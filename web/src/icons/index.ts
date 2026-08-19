import SvgIcon from '@/components/SvgIcon/index.vue'
import { App } from 'vue'

// 将目录下所有 svg 作为 symbol 注入到页面 <svg> sprite，供 <use href="#icon-xxx"> 引用。
// 原项目使用 Webpack 的 require.context + svg-sprite-loader，迁移到 Vite 后改用 import.meta.glob。
const svgModules = import.meta.glob('./svg/*.svg', {
  query: '?raw',
  import: 'default',
  eager: true
}) as Record<string, string>

function installSvgSprite() {
  if (document.getElementById('svg-sprite-container')) return
  const symbols: string[] = []
  Object.keys(svgModules).forEach((path) => {
    const name = path.split('/').pop()!.replace(/\.svg$/, '')
    const raw = svgModules[path]
    const symbol = raw
      .replace(/<svg([^>]*)>/i, `<symbol id="icon-${name}"$1>`)
      .replace(/<\/svg>/i, '</symbol>')
    symbols.push(symbol)
  })
  const container = document.createElementNS('http://www.w3.org/2000/svg', 'svg')
  container.id = 'svg-sprite-container'
  container.setAttribute('xmlns', 'http://www.w3.org/2000/svg')
  container.setAttribute('style', 'position:absolute;width:0;height:0;overflow:hidden;')
  container.innerHTML = symbols.join('')
  document.body.appendChild(container)
}

export default (app: App) => {
  installSvgSprite()
  app.component('svg-icon', SvgIcon)
}
