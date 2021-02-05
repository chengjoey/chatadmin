import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import router from './router'
import ApiService from "@/common/api.service";
import store from "./store"
import VueSSE from 'vue-sse';
import ChatWindow from 'vue-advanced-chat'
// import { IS_AUTHENTICATED} from './store/actions.type'
import jwtService from './common/jwt.service'
import Chat from 'vue-beautiful-chat'
Vue.use(Chat)

Vue.config.productionTip = false;

Vue.use(VueSSE);
Vue.use(ChatWindow);

ApiService.init();

router.beforeEach((to, from, next) => {
  // console.log(IS_AUTHENTICATED())
  if (to.name !== 'login' && !jwtService.getToken()) next({ path: '/login' })
  else next()
})

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
