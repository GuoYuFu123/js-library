import { defineComponent, reactive, toRefs } from '@vue/composition-api';
import MScrollObserve2 from './scrollObserve2.vue';
import MScrollObserve3 from './scrollObserve2.vue';

export default defineComponent({
  components: {
    MScrollObserve2,
    MScrollObserve3,
  },

  setup() {
    const state = reactive({
      list: [],
    });
    const getList = () => {
      console.log('getList');
      const arr = [1, 2, 3, 3, 4, 4, 5, 5, 5, 6, 6, 6, 66, 6];
      state.list = [...state.list, ...arr];
    };
    return {
      ...toRefs(state),
      getList,
    };
  },
  render(this) {
    return (
      <div>
        <div>{JSON.stringify(this.list)}</div>
        {this.list.map((it) => (
          <div style="height: 200px; width: 200px; margin-top:20px; border:1px solid red"></div>
        ))}

        <m-scroll-observe2
          loaderMethod={this.getList}
          loaderDisable={false}
        ></m-scroll-observe2>
        <m-scroll-observe3
          loaderMethod={this.getList}
          loaderDisable={false}
        ></m-scroll-observe3>
      </div>
    );
  },
});
