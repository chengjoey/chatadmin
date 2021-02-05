import ApiService from "@/common/api.service";
import JwtService from "@/common/jwt.service";
import jwtService from "../common/jwt.service";
import {
    CHECK_AUTH,
    LOGIN,
    LOGOUT,
} from "./actions.type";
import { PURGE_AUTH, SET_AUTH, SET_ERROR } from "./mutations.type";

const state = {
    errors: null,
    user: {},
    isAuthenticated: !!JwtService.getToken()
};

const getters = {
    currentUser(state) {
        return state.user;
    },
    isAuthenticated(state) {
        return state.isAuthenticated;
    },
    isAdmin(state) {
        return state.user.role === "admin"
    }
};

const actions = {
    [LOGIN](context, nickname) {
        return new Promise(resolve => {
            ApiService.post("login", { nickname: nickname })
                .then((response) => {
                    response.data.value.token = response.headers.token
                    context.commit(SET_AUTH, response.data.value);
                    resolve(response.data);
                })
                .catch((response) => {
                    console.log(response)
                    context.commit(SET_ERROR, response.data.msg);
                })
        });
    },
    [CHECK_AUTH](context) {
        if (jwtService.getToken()) {
            console.log("already login")
        } else {
            context.commit(PURGE_AUTH)
        }
    },
    [LOGOUT](context) {
        context.commit(PURGE_AUTH);
    }
}

const mutations = {
    [SET_ERROR](state, error) {
        state.errors = error;
    },
    [SET_AUTH](state, user) {
        state.isAuthenticated = true;
        state.user = user;
        state.errors = {},
        JwtService.saveToken(user.token);
        jwtService.setRole(user.role);
        console.log(user.role)
    },
    [PURGE_AUTH](state) {
        state.isAuthenticated = false
        state.user = {};
        state.errors = {};
        jwtService.destroyToken();
    }
}

export default {
    state,
    actions,
    mutations,
    getters,
}