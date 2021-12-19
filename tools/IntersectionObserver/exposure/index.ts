import Vue from 'vue';
import { ObserveExposure } from './exposure';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function install(Vue: any) {
  Vue.directive('guoguo-exposure', ObserveExposure);
}

const plugin = {
  version: '0.0.1',
  install,
};

Vue.use(plugin);
