<template>
  <el-container
    class="layout-container"
    :class="{ hideSidebar: !appStore.siderType }"
  >
    <el-aside :width="appStore.siderType ? '230px' : '64px'">
      <SideBar />
    </el-aside>
    <el-container class="main-container">
      <el-header>
        <NavBar />
      </el-header>
      <el-main>
        <el-scrollbar>
          <router-view v-slot="{ Component }">
            <transition name="fade-transform" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup lang="ts">
import SideBar from './components/SideBar/index.vue'
import NavBar from './components/NavBar/index.vue'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()
</script>

<style lang="scss" scoped>
.layout-container {
  position: relative;
  width: 100%;
  height: 100%;
}

// 消除导航栏内边距
.el-header {
  position: relative;
  padding: 0 !important;
  height: 48px !important;
}

// 消除主内容内边距
.el-main {
  padding: 0 !important;
  background-color: #f0f2f5;
}

.el-scrollbar {
  background-color: rgb(246, 246, 246) !important;
  padding: 10px;
}
</style>
