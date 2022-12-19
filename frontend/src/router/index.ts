import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ElecTemp from '@/components/ElecTemp.vue'
import WaterRainVue from '@/components/WaterRain.vue'
import ElecRainVue from '@/components/ElecRain.vue'
import WaterTempVue from '@/components/WaterTemp.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/elect-temp',
      name: 'elecTamp',
      component: ElecTemp
    },
    {
      path: '/elect-rainfall',
      name: 'elecRain',
      component: ElecRainVue
    },
    {
      path: '/water-temp',
      name: 'waterTemp',
      component: WaterTempVue
    },
    {
      path: '/water-rainfall',
      name: 'waterRainfall',
      component: WaterRainVue
    },
  ]
})

export default router
