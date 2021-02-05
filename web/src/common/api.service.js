import Vue from "vue";
import axios from "axios";
import VueAxios from "vue-axios";
import JwtService from "@/common/jwt.service";
import { API_URL } from "@/common/config"
import jwtService from "./jwt.service";

const ApiService = {
    init() {
        Vue.use(VueAxios, axios);
        Vue.axios.defaults.baseURL = API_URL;
    },
    getIDToken() {
        return jwtService.getToken();
    },
    setHeader() {
        Vue.axios.defaults.headers.common[
            "Authorization"
        ] = JwtService.getToken();
    },
    post(resource, params) {
        return Vue.axios.post(`${resource}`, params);
    },
    get(resource) {
        const config = {
            headers: {
                "Authorization": jwtService.getToken()
            }
        }
        return Vue.axios.get(`${resource}`, config);
    },
    put(resource, params) {
        return Vue.axios.put(`{resource}`, params)
    }
};

export default ApiService;