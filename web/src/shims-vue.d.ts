/* eslint-disable */
declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

declare module '*.svg' {
  const content: string
  export default content
}

// 副作用导入的样式文件声明（CSS / SCSS / CSS Modules）
declare module '*.css'
declare module '*.scss'
declare module '*.module.scss' {
  const classes: Record<string, string>
  export default classes
}
