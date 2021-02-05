import Vue from "vue";
import Router from "vue-router";
import Dashboard from "@/components/Dashboard"
import Chat from "@/components/Chat"

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/login",
      component: () => import("@/components/L"),
      name: "login"
    },
    {
        path: "/app",
        component: () => import("@/components/Layout"),
        name: "layout",
        children: [
          {
            path: 'dashboard',
            name: 'dashboard',
            component: Dashboard,
          },
          {
            path: "chat",
            component: Chat,
            name: "chat",
            // beforeRouteLeave() {
            //   console.log("触发离开事件");
            //   this.closewebsocket();
            // }
          }
        ]
    },
    // {
    //   path: "/",
    //   component: () => import("@/views/Home"),
    //   children: [
    //     {
    //       path: "",
    //       name: "home",
    //       component: () => import("@/views/HomeGlobal")
    //     },
    //     {
    //       path: "my-feed",
    //       name: "home-my-feed",
    //       component: () => import("@/views/HomeMyFeed")
    //     },
    //     {
    //       path: "tag/:tag",
    //       name: "home-tag",
    //       component: () => import("@/views/HomeTag")
    //     }
    //   ]
    // },
    // {
    //   name: "login",
    //   path: "/login",
    //   component: () => import("@/views/Login")
    // },
    // {
    //   name: "register",
    //   path: "/register",
    //   component: () => import("@/views/Register")
    // },
    // {
    //   name: "settings",
    //   path: "/settings",
    //   component: () => import("@/views/Settings")
    // },
    // // Handle child routes with a default, by giving the name to the
    // // child.
    // // SO: https://github.com/vuejs/vue-router/issues/777
    // {
    //   path: "/@:username",
    //   component: () => import("@/views/Profile"),
    //   children: [
    //     {
    //       path: "",
    //       name: "profile",
    //       component: () => import("@/views/ProfileArticles")
    //     },
    //     {
    //       name: "profile-favorites",
    //       path: "favorites",
    //       component: () => import("@/views/ProfileFavorited")
    //     }
    //   ]
    // },
    // {
    //   name: "article",
    //   path: "/articles/:slug",
    //   component: () => import("@/views/Article"),
    //   props: true
    // },
    // {
    //   name: "article-edit",
    //   path: "/editor/:slug?",
    //   props: true,
    //   component: () => import("@/views/ArticleEdit")
    // }
  ]
});
