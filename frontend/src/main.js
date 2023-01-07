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
            console.info("initRoutes.............................")
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
                    console.info("menus......",menus)
                    storeMenus.push.apply(storeMenus, menus.menus)
                    // 添加路由
                    this.__addRouters(storeMenus)
                    // 404
                    router.addRoute({
                        path: '*',
                        redirect: '/not-found'
                    })
                    // 路由加载完成后，如果访问的是/，跳转至动态识别的首页
                    if (this.$route.path === '/') {
                        console.info("storeMenus[0]",storeMenus[0])
                        this.$router.push(storeMenus[0].url)
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
                if (route.url == null || route.url === '') {
                    this.__addRouters(route.children, parentsDump)
                    continue
                }
                if (rs.findIndex(r => r.url === route.url) > -1) {
                    this.__addRouters(route.children, parentsDump)
                    continue
                }
                if (this.homePage == null) {
                    this.setHomePage(route)
                }
                router.addRoute('layout', {
                    path: route.url,
                    name: route.name,
                    meta: {
                        title: route.name,
                        paths: [...parents.map(p => p.name), route.name]
                    },
                    component: () => import('@/views' + route.url)
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
