import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import pinia from './stores';
import SvgIcon from '@/icons';
import './styles/index.scss';
import 'element-plus/dist/index.css';
import '@icon-park/vue-next/styles/index.css';
// 注意：Tailwind 应通过项目构建链（PostCSS + tailwind.config.js 的 content 扫描）生成，
// 直接引入 'tailwindcss/tailwind.css' 不会应用自定义主题，建议改为项目内的 Tailwind 入口文件。
import 'tailwindcss/tailwind.css';
import './styles/theme.scss';
import { MotionPlugin } from '@vueuse/motion';
import ElementPlus from 'element-plus';

const app = createApp(App);

// 全局错误兜底，避免生产环境白屏难排查
app.config.errorHandler = (err: unknown, _vm: unknown, info: string) => {
  // eslint-disable-next-line no-console
  console.error('[Global Error]', err, info);
};

// 挂载 Pinia 状态管理（替代原 Vuex）
app.use(pinia);

SvgIcon(app);

app
  .use(router)
  .use(ElementPlus)
  .use(MotionPlugin);

// 等待路由（含异步组件/导航守卫）就绪后再挂载，避免首屏闪烁或路由未解析
router.isReady().then(() => {
  app.mount('#app');
});
