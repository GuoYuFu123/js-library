<template>
  <div v-if="!loaderDisable" ref="scrollLoader">loading...</div>
</template>

<script lang="ts">
import {
  defineComponent,
  ref,
  computed,
  onMounted,
  onActivated,
  onDeactivated,
  onBeforeUnmount,
} from '@vue/composition-api';
type Props = {
  loaderMethod: VoidFunction;
  loaderViewport: Element;
  loaderDistance: number;
  loaderDisable: boolean;
};
export default defineComponent<Props>({
  props: {
    loaderMethod: {
      default: () => {},
      required: true,
    },
    loaderViewport: {
      type: Element,
      default: null,
    },
    loaderDistance: {
      type: Number,
      default: 0,
    },
    loaderDisable: {
      type: Boolean,
      default: false,
    },
  },
  setup(props) {
    const scrollLoader = ref<Element>();
    const options = computed(() => {
      return {
        root: props.loaderViewport,
        rootMargin: `0px 0px ${props.loaderDistance}px 0px`,
      };
    });
    const _observer = computed(() => {
      return new IntersectionObserver(([{ isIntersecting }]) => {
        isIntersecting && !props.loaderDisable && props.loaderMethod();
      }, options.value);
    });
    onMounted(() => {
      if (scrollLoader?.value) _observer.value.observe(scrollLoader?.value);
    });
    onActivated(() => {
      if (scrollLoader?.value) _observer.value.observe(scrollLoader?.value);
    });
    onDeactivated(() => {
      if (scrollLoader?.value) _observer.value.unobserve(scrollLoader?.value);
    });
    onBeforeUnmount(() => {
      if (scrollLoader?.value) _observer.value.unobserve(scrollLoader?.value);
    });
    return { _observer, scrollLoader };
  },
});
</script>
