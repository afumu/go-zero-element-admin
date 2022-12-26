import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import ElementUI from 'element-ui';
import './assets/style/element-variables.scss';
import VueClipboard from 'vue-clipboard2';
import directives from './directives';
import filters from './filters';
import plugins from './plugins';
import 'default-passive-events';
import Print from 'vue-print-nb';
import {
    mapState,
    mapMutations
} from 'vuex'
// import { fetchMenuTree } from './api/system/menu'
import {
    getMenu
} from './api/system/common'
import VueQuillEditor from 'vue-quill-editor'
import VueCron from 'vue-cron'

Vue.config.productionTip = false
Vue.use(ElementUI, {
    size: 'small'
})
Vue.use(VueClipboard);
Vue.use(directives);
Vue.use(filters);
Vue.use(plugins);
Vue.use(VueCron);
Vue.use(VueQuillEditor);
Vue.use(Print);

new Vue({
    data: {
        loading: false
    },
    router,
    store,
    computed: {
        ...mapState(['userInfo', 'homePage'])
    },
    watch: {
        async userInfo() {
            if (this.userInfo == null) {
                return
            }
            await this.initRoutes()
        }
    },
    methods: {
        ...mapMutations(['switchCollapseMenu', 'setHomePage']),
        // 初始化本地配置
        initLocalConfig() {

            // 菜单状态配置
            const menuStatus = window.localStorage.getItem('MENU_STATUS')
            if (menuStatus != null) {
                this.switchCollapseMenu(menuStatus === 'true')
            }
        },
        // 初始化路由
        async initRoutes() {
            if (this.loading || this.userInfo == null) {
                return
            }
            this.loading = true
            // 重置菜单
            this.$store.commit('resetMenus')
            // 获取菜单
            const storeMenus = this.$store.state.menuData.list
            if (storeMenus.length > 0 && this.homePage == null) {
                this.setHomePage(storeMenus[0])
            }
            await getMenu()
                .then(menus => {
                    // this.$store.commit('SET_MENU', menus.menu)
                    // console.log('menuData', this.menuData)
                    // 添加菜单
                    storeMenus.push.apply(storeMenus, menus.menu)
                    // 添加路由
                    this.__addRouters(storeMenus)
                    // 404
                    router.addRoute({
                        path: '*',
                        redirect: '/not-found'
                    })
                    // 路由加载完成后，如果访问的是/，跳转至动态识别的首页
                    if (this.$route.path === '/') {
                        this.$router.push(storeMenus[0].path)
                    }
                })
                .catch(e => {
                    throw e
                })
                .finally(() => {
                    this.loading = false
                })
        },
        // 新建路由
        __addRouters(routes, parents = []) {
            if (routes == null || routes.length === 0) {
                return
            }
            const rs = router.getRoutes()
            // console.log('routes', routes)
            for (const route of routes) {
                const parentsDump = JSON.parse(JSON.stringify(parents))
                parentsDump.push(route)
                if (route.path == null || route.path === '') {
                    this.__addRouters(route.children, parentsDump)
                    continue
                }
                if (rs.findIndex(r => r.path === route.path) > -1) {
                    this.__addRouters(route.children, parentsDump)
                    continue
                }
                if (this.homePage == null) {
                    this.setHomePage(route)
                }
                router.addRoute('layout', {
                    path: route.path,
                    name: route.name,
                    meta: {
                        title: route.meta.title,
                        paths: [...parents.map(p => p.meta.title), route.meta.title]
                    },
                    component: () => import('@/views' + route.path)
                })
                this.__addRouters(route.children, parentsDump)
            }
        }
    },
    async created() {
        if (this.userInfo == null) {
            return
        }
        await this.initRoutes()
            .catch(() => {})
    },
    mounted() {
        this.initLocalConfig()
    },
    render: h => h(App)
}).$mount('#app')