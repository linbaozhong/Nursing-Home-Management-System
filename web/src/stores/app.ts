import { defineStore } from 'pinia'
import { IRoute } from '@/router/types'
import { getRouterList } from '@/mock/getters'
import { clearRoutes, formatRouteTree } from '@/router/utils'
// import { getUserRouteList } from '@/apis'
import { getLogin } from '@/apis/user'
import router from '@/router'

interface IAppState {
  uid: number
  token: string
  hasAuth: boolean
  routeTree: IRoute[]
  siderType: boolean
  userPeofile: any
  rememberPWD: boolean
}

const defaultState: IAppState = {
  uid: 1,
  token: '',
  hasAuth: false,
  routeTree: [],
  siderType: true,
  userPeofile: {},
  rememberPWD: false
}

// 替代原 Vuex 的 app 模块（namespaced: app）
export const useAppStore = defineStore('app', {
  state: (): IAppState => ({ ...defaultState }),
  actions: {
    // 登录
    async actionLogin(data: any) {
      const res: any = await getLogin(data)
      if (res.code === 200) {
        this.token = res?.data?.token
        this.userPeofile = {
          username: res?.data.name,
          userid: res?.data.id,
          avator: res?.data.avator,
          authIdList: res?.data.authIdList
        }
      }
      this.rememberPWD = data?.rememberPWD
      return res
    },
    // 获取权限菜单
    getRouterTree() {
      // 模拟数据
      const routeList = getRouterList(this.uid).data as unknown as IRoute[]
      this.routeTree = formatRouteTree(routeList)
      this.hasAuth = true
    },
    // 点击登出
    logout() {
      this.$reset()
      localStorage.clear()
      clearRoutes()
      router.replace('/login')
    }
  }
})
