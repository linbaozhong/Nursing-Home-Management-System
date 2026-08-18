<template>
  <el-config-provider :locale="locale">
    <div class="app">
      <!-- 渲染期异常兜底，避免单个页面错误导致整站白屏 -->
      <div v-if="renderError" class="app-error">
        <p>页面渲染出现异常，请刷新重试。</p>
        <pre v-if="errorText">{{ errorText }}</pre>
        <button class="app-error__retry" type="button" @click="retry">重试</button>
      </div>
      <router-view v-else />
    </div>
  </el-config-provider>
</template>
<script setup lang="ts">
import  ElConfigProvider  from 'element-plus'
import { computed, onErrorCaptured, ref } from 'vue'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

const locale = zhCn

// 顶层错误边界：捕获子树渲染期错误，避免整站白屏
const renderError = ref<unknown>(null)
onErrorCaptured((err: unknown) => {
  renderError.value = err
  // 返回 false 阻止错误继续向上冒泡导致整棵组件树卸载
  return false
})

const errorText = computed(() =>
  renderError.value instanceof Error ? renderError.value.message : String(renderError.value)
)

// 重试：清除错误状态，重新渲染路由视图（处理瞬时渲染异常）
function retry() {
  renderError.value = null
}
</script>

<style lang="scss" scoped>
.app {
  width: 100%;
  height: 100%;
}

.app-error {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #f56c6c;
  font-size: 14px;

  pre {
    margin-top: 8px;
    max-width: 90%;
    padding: 8px 12px;
    overflow: auto;
    color: #909399;
    background: #f5f7fa;
    border-radius: 4px;
  }

  &__retry {
    margin-top: 16px;
    padding: 8px 20px;
    color: #fff;
    background: #409eff;
    border: none;
    border-radius: 4px;
    cursor: pointer;

    &:hover {
      background: #66b1ff;
    }
  }
}
</style>
