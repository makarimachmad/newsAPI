import Vue from 'vue'
import Router from 'vue-router'


import Beranda from './components/Beranda.vue'
import DetailBisnis from './components/DetailBisnis.vue'
import DetailHealth from './components/DetailHealth.vue'
import DetailEntertainment from './components/DetailEntertainment.vue'



Vue.use(Router)

export default new Router ({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'beranda',
            component: Beranda
        },
        {
            path: '/detailbisnis/:id',
            name: 'detailbisnis',
            component: DetailBisnis
        },
        {
            path: '/detailhealth/:id',
            name: 'detailhealth',
            component: DetailEntertainment
        }
        ,
        {
            path: '/detailentertainment/:id',
            name: 'detailentertainment',
            component: DetailHealth
        }
    ]
})