<template>
  <div class="container">
    <el-tooltip effect="light" :content="content" placement="right">
      <svg-icon
        icon="menu-fold"
        @click="changeSiderType"
        :style="style"
      ></svg-icon>
    </el-tooltip>
  </div>
</template>

<script setup lang="ts">
import { useAppStore } from '@/stores/app'
import { computed } from 'vue'

const appStore = useAppStore()

const content = computed(() => {
  return appStore.siderType ? '点击折叠' : '点击展开'
})

const style = computed(() => {
  return appStore.siderType ? 'none' : 'transform: rotateY(180deg);'
})

const changeSiderType = () => {
  appStore.siderType = !appStore.siderType
}
</script>

<style lang="scss" scoped>
.svg-icon {
  color: var(--el-color-primary);
  margin: 0 0 4px 16px !important;
  vertical-align: middle;
  outline: none;
  transition-duration: 0.36s;
  cursor: pointer;
}
</style>
