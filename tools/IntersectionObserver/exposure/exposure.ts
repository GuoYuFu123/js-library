import 'intersection-observer';
import { analytics } from '@guoguo/analytics';
import { VNode } from 'vue/types/umd';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
(IntersectionObserver.prototype as any)['THROTTLE_TIMEOUT'] = 300;

const intersectionObserverOptions = {
  root: null,
  rootMargin: '0px',
  threshold: 0.2, // 不一定非得全部露出来  这个阈值可以小一点点
};

interface IValue {
  name: string;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  params: any;
}
export class Exposure {
  private el: HTMLElement;
  private value: IValue;
  private _observer: IntersectionObserver | null;
  constructor(el: HTMLElement, value: IValue, vnode: VNode) {
    this.el = el;
    this.value = value;
    this._observer = null;

    this.register(value, vnode);
  }
  register(value: IValue, vnode: VNode) {
    if (this._observer) {
      this.unRegister();
    }
    this._observer = new IntersectionObserver(entries => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          const { name, params } = value;
          if (params) {
            analytics.track(name, params);
          } else {
            analytics.track(name);
          }

          this.unRegister();
        }
      });
    }, intersectionObserverOptions);
    // 进行dom节点监听, 因为存在dom节点一开始就在可视区，有数据后把该节点挤到页面下边
    setTimeout(() => {
      vnode &&
        vnode.context?.$nextTick(() => {
          this._observer?.observe(this.el);
        });
    }, 200);
  }
  unRegister() {
    if (this._observer) {
      this._observer.disconnect();
      this._observer = null;
    }
  }
}
export const ObserveExposure = {
  bind: (el: HTMLElement, { value }: { value: IValue }, vnode: VNode) => {
    if (!value) {
      return false;
    }

    const v = value as IValue;
    new Exposure(el, v, vnode);
  },
};
