(function () {
    // 可选数据源（用于下拉搜索），按 input 的 data-source 区分
    var SOURCES = {
        market: [
            'Germany (DE)', 'France (FR)', 'Italy (IT)', 'Spain (ES)',
            'United Kingdom (UK)', 'Netherlands (NL)', 'Belgium (BE)',
            'Poland (PL)', 'Sweden (SE)', 'Denmark (DK)', 'Finland (FI)',
            'Norway (NO)', 'Switzerland (CH)', 'Austria (AT)', 'Portugal (PT)',
            'Ireland (IE)', 'Czechia (CZ)', 'Romania (RO)', 'Greece (GR)'
        ],
        gender: ['Female', 'Male', 'Kids'],
        placement: [
            'Meta Story / Link / Reels',
            'Pinterest Static Ad',
            'TikTok Top Feed / In-Feed / Spark',
            'YouTube Video Reach / View',
            'Meta Collection / Carousel',
            'TikTok Video Shopping Ad / Smart+'
        ],
        'placement-aps': [
            'Awareness Teaser App',
            'Awareness Teaser Web',
            'Sponsored Product Ads',
            'Collection Teaser Catalog App',
            'Collection Teaser Catalog Web',
            'Collection Teaser Homepage App',
            'Collection Teaser Homepage Video',
            'Collection Teaser Homepage Web',
            'In Catalog App',
            'In Catalog Web'
        ],
        account: [
            'Adidas AG', 'Nike Inc.', 'PUMA SE', 'Under Armour Inc.',
            'Lululemon Athletica', 'ASICS Corp.', 'New Balance', 'Skechers USA',
            'Anta Sports', 'Li-Ning Co.', 'Xtep International', 'Fila (Korea)'
        ],
        // 各服饰公司对应的全部品牌（用于 Brands 联动）
        brandByAccount: {
            'Adidas AG': ['Adidas', 'Adidas Originals', 'Adidas Performance', 'Reebok', 'TaylorMade', 'Salomon', 'Wilson', 'CCM Hockey'],
            'Nike Inc.': ['Nike', 'Air Jordan', 'Converse', 'Hurley'],
            'PUMA SE': ['PUMA', 'Cobra Golf'],
            'Under Armour Inc.': ['Under Armour', 'MapMyRun'],
            'Lululemon Athletica': ['Lululemon'],
            'ASICS Corp.': ['ASICS', 'Onitsuka Tiger'],
            'New Balance': ['New Balance'],
            'Skechers USA': ['Skechers'],
            'Anta Sports': ['Anta', 'FILA China', 'Salomon (Anta)'],
            'Li-Ning Co.': ['Li-Ning'],
            'Xtep International': ['Xtep'],
            'Fila (Korea)': ['FILA']
        }
    };

    // 未指定 data-source 时默认使用市场数据源
    function getOptions(input) {
        var key = input.getAttribute('data-source') || 'market';
        if (key === 'brand') {
            // 联动：根据关联的 Account 当前选中值返回对应品牌列表
            var accInput = document.querySelector('input[data-source="account"]');
            var accName = accInput ? (accInput.value || '').trim() : '';
            return SOURCES.brandByAccount[accName] || [];
        }
        return SOURCES[key] || SOURCES.market;
    }

    // Account 改选后，清空各 Brands 列表已选品牌并刷新下拉选项
    function syncBrandOptions() {
        var brandInputs = document.querySelectorAll('input[data-source="brand"]');
        Array.prototype.forEach.call(brandInputs, function (bInput) {
            var list = bInput.closest('.market-list');
            if (list) {
                Array.prototype.forEach.call(list.querySelectorAll('.market-tag'), function (tag) {
                    tag.remove();
                });
                updatePlacementHint(list);
            }
            bInput.dispatchEvent(new Event('input')); // 触发对应 combo 的 render 重算
        });
    }

    // 全局事件委托：处理所有 .market-tag 内 × 按钮的删除（含原有与新增）
    document.addEventListener('click', function (e) {
        var btn = e.target.closest('.market-tag button');
        if (btn) {
            var tag = btn.closest('.market-tag');
            if (tag) {
                var list = tag.closest('.market-list');
                tag.remove();
                if (list) {
                    updatePlacementHint(list);
                    if (list.getAttribute('data-gender') === '1') {
                        updateGenderTotal(list);
                        refreshGenderPlaceholders(list);
                    }
                    // 删除后可能有可选项重新出现，触发对应搜索框的 render 恢复显示
                    var si = list.querySelector('input.add-input');
                    if (si) si.dispatchEvent(new Event('input'));
                }
            }
        }
    });

    // 根据列表内是否有选中项，切换其前方 .placement-hint 的显隐
    function updatePlacementHint(list) {
        var node = list;
        var hint = null;
        // 向上查找：list 之前的兄弟 .placement-hint，或祖先内 list 之前的 .placement-hint
        while (node && !hint) {
            var prev = node.previousElementSibling;
            while (prev) {
                if (prev.classList && prev.classList.contains('placement-hint')) { hint = prev; break; }
                prev = prev.previousElementSibling;
            }
            node = node.parentElement;
        }
        if (hint) {
            var hasTag = list.querySelector('.market-tag');
            hint.style.display = hasTag ? 'none' : '';
        }
    }

    function tagText(tag) {
        var text = '';
        tag.childNodes.forEach(function (n) {
            if (n.nodeType === 3) text += n.textContent;
        });
        return text.trim();
    }

    function tagExists(list, text) {
        return Array.prototype.some.call(list.querySelectorAll('.market-tag'), function (t) {
            return tagText(t) === text;
        });
    }

    function appendTag(list, text) {
        if (tagExists(list, text)) return;
        var tag = document.createElement('span');
        tag.className = 'market-tag';
        tag.appendChild(document.createTextNode(text + ' '));
        var btn = document.createElement('button');
        btn.textContent = '×';
        tag.appendChild(btn);
        // 若属于 gender 列表，附加一个百分比输入框
        if (list.getAttribute('data-gender') === '1') {
            var pct = document.createElement('input');
            pct.type = 'text';
            pct.className = 'gender-pct';
            pct.placeholder = '%';
            pct.setAttribute('inputmode', 'decimal');
            bindGenderPct(list, pct);
            tag.appendChild(pct);
        }
        // 插入到列表内搜索框之前（让搜索框始终在末尾）；无搜索框则追加到末尾
        // 注意：必须用 .add-input 限定，避免选中 tag 内部的 .gender-pct 等输入框
        var input = list.querySelector('input.add-input');
        if (input) {
            list.insertBefore(tag, input);
        } else {
            list.appendChild(tag);
        }
        updatePlacementHint(list);
        if (list.getAttribute('data-gender') === '1') updateGenderTotal(list);
    }

    // 计算并更新 gender 列表的合计行（合计元素由 HTML 静态提供，不再动态新建）
    function updateGenderTotal(list) {
        var wrap = list.closest('.market-col') || list.parentElement;
        var totalEl = list.parentElement.querySelector('.gender-total') ||
            (wrap && wrap !== list.parentElement ? wrap.querySelector('.gender-total') : null);
        if (!totalEl) return; // 页面未提供合计元素则不处理
        var sum = 0;
        Array.prototype.forEach.call(list.querySelectorAll('.gender-pct'), function (p) {
            var v = parseFloat((p.value || '').replace('%', ''));
            if (!isNaN(v)) sum += v;
        });
        totalEl.textContent = sum + '%';
        totalEl.classList.toggle('warn', sum !== 100);
    }

    // 绑定 gender 百分比框的 input 事件：总数超过 100 时本次输入无效（恢复原值）
    function bindGenderPct(list, pct) {
        if (pct._genderBound) return;
        pct._genderBound = true;
        pct._last = pct.value || '';
        pct.addEventListener('input', function () {
            var raw = (pct.value || '').trim().replace('%', '');
            var cur = parseFloat(raw);
            if (raw !== '' && isNaN(cur)) {
                // 非数字输入：恢复原值
                pct.value = pct._last;
                return;
            }
            // 计算其它框已填之和
            var others = 0;
            Array.prototype.forEach.call(list.querySelectorAll('.gender-pct'), function (p) {
                if (p === pct) return;
                var v = parseFloat((p.value || '').replace('%', ''));
                if (!isNaN(v)) others += v;
            });
            if (raw !== '' && others + cur > 100) {
                // 超过 100：本次输入无效，恢复为上次有效值
                pct.value = pct._last;
                return;
            }
            pct._last = pct.value;
            updateGenderTotal(list);
            refreshGenderPlaceholders(list);
        });
    }

    // 根据已填总和，将同列表中仍为空白的百分比框 placeholder 置为剩余量
    function refreshGenderPlaceholders(list) {
        var sum = 0;
        Array.prototype.forEach.call(list.querySelectorAll('.gender-pct'), function (p) {
            var v = parseFloat((p.value || '').replace('%', ''));
            if (!isNaN(v)) sum += v;
        });
        var remain = Math.max(0, 100 - sum);
        Array.prototype.forEach.call(list.querySelectorAll('.gender-pct'), function (p) {
            if ((p.value || '').trim() === '') {
                p.placeholder = remain > 0 ? remain + '%' : '0%';
            } else {
                p.placeholder = '%';
            }
        });
    }

    // 找到位于 combo 之前、且最近的一个 .market-list（DOM 顺序上的上一个列表）
    function findTargetList(combo) {
        var node = combo.parentElement;
        var guard = 0;
        while (node && guard++ < 20) {
            var children = node.children;
            var found = null;
            for (var i = 0; i < children.length; i++) {
                if (children[i] === combo) break; // 只在该 combo 之前的兄弟里找
                if (children[i].classList && children[i].classList.contains('market-list')) {
                    found = children[i]; // 取最后一个（最近的）
                }
            }
            if (found) return found;
            node = node.parentElement;
        }
        var list = document.createElement('div');
        list.className = 'market-list';
        combo.parentElement.insertBefore(list, combo);
        return list;
    }

    function closeDropdown(combo) {
        combo.classList.remove('open');
        document.removeEventListener('click', onDocClick, true);
    }

    function onDocClick(e) {
        var combo = document.querySelector('.add-combo.open');
        if (combo && !combo.contains(e.target)) {
            closeDropdown(combo);
        }
    }

    // 绑定一个 combo：container 为 .add-combo 元素，input 为搜索框，list 为目标列表
    function bindCombo(container, input, list) {
        container.classList.add('add-combo');

        var dropdown = document.createElement('div');
        dropdown.className = 'add-dropdown';
        container.appendChild(dropdown);

        var activeIndex = -1;
        var currentOptions = [];

        var isSingle = input.getAttribute('data-single') === '1';

        function render() {
            var q = input.value.trim().toLowerCase();
            var source = getOptions(input);
            var isGender = (input.getAttribute('data-source') || 'market') === 'gender';
            // 单选模式：始终展示完整源（可重新选择），不过滤已选项
            currentOptions = source.filter(function (m) {
                if (isSingle) return m.toLowerCase().indexOf(q) !== -1;
                return m.toLowerCase().indexOf(q) !== -1 && !tagExists(list, m);
            });
            dropdown.innerHTML = '';
            activeIndex = -1;
            if (currentOptions.length === 0) {
                var empty = document.createElement('div');
                empty.className = 'add-option empty';
                var emptyText = isGender ? 'No matching gender'
                    : isSingle ? 'No matching account'
                    : 'No matching market';
                empty.textContent = emptyText;
                dropdown.appendChild(empty);
            } else {
                currentOptions.forEach(function (m) {
                    var opt = document.createElement('div');
                    opt.className = 'add-option';
                    opt.dataset.value = m;
                    opt.textContent = m;
                    dropdown.appendChild(opt);
                });
            }
            if (isSingle) return; // 单选模式不隐藏输入框
            // 当所有可选项都已被选（无剩余可选项）时，仅隐藏 input 与下拉框；否则显示
            var hasRemaining = source.some(function (m) { return !tagExists(list, m); });
            input.style.display = hasRemaining ? '' : 'none';
            dropdown.style.display = hasRemaining ? '' : 'none';
            container.classList.toggle('empty-source', !hasRemaining);
        }

        dropdown.addEventListener('click', function (e) {
            var opt = e.target.closest('.add-option:not(.empty)');
            if (!opt) return;
            e.stopPropagation();
            if (isSingle) {
                input.value = opt.dataset.value;   // 单选：填入搜索框
                closeDropdown(container);
                if (input.getAttribute('data-source') === 'account') syncBrandOptions();
                return;
            }
            appendTag(list, opt.dataset.value);
            render();           // 已选项从下拉消失，但保留搜索框内容
            input.focus();      // 保持焦点，下拉不关闭；失焦时才清空并隐藏
        });

        function open() {
            var other = document.querySelector('.add-combo.open');
            if (other && other !== container) closeDropdown(other);
            container.classList.add('open');
            document.addEventListener('click', onDocClick, true);
            render();
        }

        input.addEventListener('focus', open);
        input.addEventListener('click', function (e) {
            e.stopPropagation();
            if (!container.classList.contains('open')) open();
        });
        input.addEventListener('input', render);
        input.addEventListener('keydown', function (e) {
            var opts = dropdown.querySelectorAll('.add-option:not(.empty)');
            if (e.key === 'ArrowDown') {
                e.preventDefault();
                if (!container.classList.contains('open')) { open(); return; }
                activeIndex = Math.min(activeIndex + 1, opts.length - 1);
            } else if (e.key === 'ArrowUp') {
                e.preventDefault();
                activeIndex = Math.max(activeIndex - 1, 0);
            } else if (e.key === 'Enter') {
                e.preventDefault();
                if (isSingle) {
                    if (currentOptions.length > 0) {
                        input.value = currentOptions[activeIndex >= 0 ? activeIndex : 0];
                        closeDropdown(container);
                        if (input.getAttribute('data-source') === 'account') syncBrandOptions();
                    }
                    return;
                }
                if (activeIndex >= 0 && opts[activeIndex]) {
                    appendTag(list, currentOptions[activeIndex]);
                } else if (currentOptions.length > 0) {
                    appendTag(list, currentOptions[0]);
                } else {
                    return;
                }
                render();           // 已选项从下拉消失，保留搜索框内容
                input.focus();      // 保持打开，失焦时才清空并隐藏
                return;
            } else if (e.key === 'Escape') {
                closeDropdown(container);
                input.blur();
                return;
            } else {
                return;
            }
            opts.forEach(function (o, i) { o.classList.toggle('active', i === activeIndex); });
        });

        // mousedown 阻止默认，避免点击选项时 input 失焦导致下拉关闭
        dropdown.addEventListener('mousedown', function (e) {
            if (e.target.closest('.add-option:not(.empty)')) e.preventDefault();
        });

        // 失去焦点后关闭下拉；多选模式清空输入框，单选模式保留选中值
        input.addEventListener('blur', function () {
            if (!isSingle) input.value = '';
            closeDropdown(container);
        });
    }

    // 初始化：
    // A) 所有 .add-input 直接绑定到其所在 .market-list（兼容你把 input 放进 .market-list 的写法）
    // B) 仍为纯文本的 .add-link（内部无 input）-> 替换成输入框（原文字作 placeholder）
    document.querySelectorAll('.add-input').forEach(function (input) {
        var list = input.closest('.market-list');
        if (!list) {
            // input 不在 .market-list 内时，找它之前最近的 .market-list
            list = findTargetList(input);
        }
        // 容器：input 最近的已有 .add-combo，否则用其直接父元素（加 add-combo 类以定位下拉）
        var container = input.closest('.add-combo') || input.parentElement;
        bindCombo(container, input, list);
    });

    document.querySelectorAll('.add-link').forEach(function (link) {
        // 仅处理仍是无 input 的纯文本写法，避免重复绑定
        if (link.querySelector('input')) return;
        var placeholder = link.textContent.trim();
        var combo = document.createElement('div');
        combo.className = 'add-combo';

        var input = document.createElement('input');
        input.type = 'text';
        input.className = 'add-input';
        input.placeholder = placeholder || 'Add market...';

        var dropdown = document.createElement('div');
        dropdown.className = 'add-dropdown';

        combo.appendChild(input);
        combo.appendChild(dropdown);
        link.parentElement.replaceChild(combo, link);

        var list = findTargetList(combo);
        bindCombo(combo, input, list);
    });

    // 页面加载时，根据各列表已有选中项刷新对应的 .placement-hint 初始显隐
    document.querySelectorAll('.market-list').forEach(function (list) {
        updatePlacementHint(list);
    });

    // 数值步进器：解析 "800,000.00 EUR" 格式，按步长增减
    (function () {
        var STEP = 1000; // 步长（与 hint 一致）
        var steppers = document.querySelectorAll('.number-stepper');

        function parseValue(text) {
            // 提取数字部分（处理千分位逗号与小数）
            var m = (text || '').match(/[\d,.]+/);
            if (!m) return 0;
            return parseFloat(m[0].replace(/,/g, '')) || 0;
        }

        function formatValue(num) {
            // 保留两位小数并加千分位
            return num.toLocaleString('en-US', {
                minimumFractionDigits: 2,
                maximumFractionDigits: 2
            });
        }

        steppers.forEach(function (stepper) {
            var input = stepper.querySelector('input');
            var up = stepper.querySelector('.step-up');
            var down = stepper.querySelector('.step-down');
            if (!input) return;

            function change(delta) {
                var current = parseValue(input.value);
                var next = current + delta * STEP;
                if (next < 0) next = 0;
                input.value = formatValue(next);
            }

            // 输入纯数字（或带逗号/小数）后，失焦时自动格式化为 #,###.00
            input.addEventListener('blur', function () {
                var current = parseValue(input.value);
                if (!isNaN(current)) input.value = formatValue(current);
            });

            if (up) up.addEventListener('click', function () { change(1); });
            if (down) down.addEventListener('click', function () { change(-1); });
        });
    })();

    // ---- 初始化：给 gender 列表中已有（HTML 静态）的百分比框绑定 input 事件并渲染合计 ----
        document.querySelectorAll('.market-list[data-gender="1"]').forEach(function (list) {
            Array.prototype.forEach.call(list.querySelectorAll('.gender-pct'), function (p) {
                bindGenderPct(list, p);
            });
            updateGenderTotal(list);
            refreshGenderPlaceholders(list);
        });

    // ---- Offsite placements 顶部的 Copy / Clean 操作 ----
    function findColByTitle(titleText) {
        var cols = document.querySelectorAll('.market-col');
        for (var i = 0; i < cols.length; i++) {
            var h = cols[i].querySelector('h3');
            if (h && h.textContent.indexOf(titleText) !== -1) return cols[i];
        }
        return null;
    }

    // 取某个 market-col 内、按 DOM 顺序的前 N 个 .market-list（即 Awareness/Engagement/Performance）
    function funnelLists(col) {
        return col ? Array.prototype.slice.call(col.querySelectorAll('.market-list'), 0, 3) : [];
    }

    // 复制 Onsite 三个相位的已选市场到 Offsite 对应相位
    function copyFromOnsite() {
        var srcCol = findColByTitle('Onsite placements');
        var dstCol = findColByTitle('Offsite placements');
        if (!srcCol || !dstCol) return;
        var srcLists = funnelLists(srcCol);
        var dstLists = funnelLists(dstCol);
        for (var i = 0; i < srcLists.length; i++) {
            var src = srcLists[i];
            var dst = dstLists[i];
            if (!dst) continue;
            Array.prototype.forEach.call(src.querySelectorAll('.market-tag'), function (tag) {
                appendTag(dst, tagText(tag));
            });
        }
    }

    // 清空 Offsite 区块内所有已选 tag（保留输入框），并刷新提示
    function cleanOffsiteAll() {
        var dstCol = findColByTitle('Offsite placements');
        if (!dstCol) return;
        Array.prototype.forEach.call(dstCol.querySelectorAll('.market-tag'), function (tag) {
            tag.remove();
        });
        Array.prototype.forEach.call(dstCol.querySelectorAll('.market-list'), function (list) {
            updatePlacementHint(list);
        });
    }

    var copyBtn = document.querySelector('.link-btn[data-action="copy-onsite"]');
    var cleanBtn = document.querySelector('.link-btn[data-action="clean-all"]');
    if (copyBtn) copyBtn.addEventListener('click', copyFromOnsite);
    if (cleanBtn) cleanBtn.addEventListener('click', cleanOffsiteAll);
})();
