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

interface ILogElement extends HTMLElement {
  logData: IValue;
}

export class Exposure {
  private _observer: IntersectionObserver | null;
  private elementMap = new Map<Element, Element>();
  constructor() {
    this.register();
  }
  register() {
    if (!this._observer) {
      this._observer = new IntersectionObserver(entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            const { name, params } = (entry.target as ILogElement).logData;
            if (params) {
              analytics.track(name, params);
            } else {
              analytics.track(name);
            }
            this.unobserve(entry.target);
          }
        });
      }, intersectionObserverOptions);
    }
  }
  // 监听
  observe(el: Element) {
    this._observer?.observe(el);
    this.elementMap.set(el, el);
  }
  // 移出监听
  unobserve(el: Element) {
    this._observer?.unobserve(el);
    this.elementMap.delete(el);

    this.disconnent();
  }
  // 断开连接
  disconnent() {
    this.elementMap.size === 0 && this._observer?.disconnect();
  }
}

// 初始化Exposure实例
const exposureInstance = new Exposure();

export const ObserveExposure = {
  bind: (el: ILogElement, { value }: { value: IValue }, vnode: VNode) => {
    if (!value) {
      return false;
    }
    setTimeout(() => {
      vnode?.context?.$nextTick(() => {
        el.logData = value;
        exposureInstance.observe(el);
      });
    }, 200);
  },
};
