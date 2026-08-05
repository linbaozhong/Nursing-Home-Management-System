(function () {
    // 可选市场数据源（用于下拉搜索）
    var MARKETS = [
        'Germany (DE)', 'France (FR)', 'Italy (IT)', 'Spain (ES)',
        'United Kingdom (UK)', 'Netherlands (NL)', 'Belgium (BE)',
        'Poland (PL)', 'Sweden (SE)', 'Denmark (DK)', 'Finland (FI)',
        'Norway (NO)', 'Switzerland (CH)', 'Austria (AT)', 'Portugal (PT)',
        'Ireland (IE)', 'Czechia (CZ)', 'Romania (RO)', 'Greece (GR)'
    ];

    // 全局事件委托：处理所有 .market-tag 内 × 按钮的删除（含原有与新增）
    document.addEventListener('click', function (e) {
        var btn = e.target.closest('.market-tag button');
        if (btn) {
            var tag = btn.closest('.market-tag');
            if (tag) {
                var list = tag.closest('.market-list');
                tag.remove();
                if (list) updatePlacementHint(list);
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
        // 插入到列表内第一个 input 之前（让搜索框始终在末尾）；无 input 则追加到末尾
        var input = list.querySelector('input');
        if (input) {
            list.insertBefore(tag, input);
        } else {
            list.appendChild(tag);
        }
        updatePlacementHint(list);
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

        function render() {
            var q = input.value.trim().toLowerCase();
            currentOptions = MARKETS.filter(function (m) {
                return m.toLowerCase().indexOf(q) !== -1 && !tagExists(list, m);
            });
            dropdown.innerHTML = '';
            activeIndex = -1;
            if (currentOptions.length === 0) {
                var empty = document.createElement('div');
                empty.className = 'add-option empty';
                empty.textContent = 'No matching market';
                dropdown.appendChild(empty);
                return;
            }
            currentOptions.forEach(function (m) {
                var opt = document.createElement('div');
                opt.className = 'add-option';
                opt.dataset.value = m;
                opt.textContent = m;
                dropdown.appendChild(opt);
            });
        }

        dropdown.addEventListener('click', function (e) {
            var opt = e.target.closest('.add-option:not(.empty)');
            if (!opt) return;
            e.stopPropagation();
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

        // 失去焦点后清空输入框并关闭下拉
        input.addEventListener('blur', function () {
            input.value = '';
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
})();
