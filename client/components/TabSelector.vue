<template>
      <ul class="flex rounded">
        <li v-for="tab in tabs" class="mr-3">
          <a
          class="inline-block border rounded py-1 px-3"
          v-bind:class="{ 'border-gray-700 bg-gray-700 text-white': tab.id === currentTabId, 'border-white hover:border-gray-200 hover:bg-gray-200': tab.id !== currentTabId}"
          href="#"
          @click="clickTab(tab.id)"
          >{{ tab.title }}</a>
        </li>
      </ul>
</template>

<script>
export default {
  data() {
    return {
      tabs: [
          {id: 'system', title: 'System'},
        ],
      currentTabId: 'system',
    }
  },
  mounted() {
    this.$root.$on('selectUser', user => {
      console.log('user selected catched');

      let contains = false;

      for(let i = 0; i < this.tabs.length; i++) {
        if (this.tabs[i].id === user.id) {
          contains = true;
          break;
        }
      }

      if (!contains) {
        this.createNewTabForUser(user);
      } else {
        if (this.currentTabId !== user.id) {
        this.currentTabId = user.id;
        this.$root.$emit('selectedTabChanged', user.id);
      }
      }
    });
  },
  methods: {
    removeTab() {

    },
    clickTab(id) {
      if (this.currentTabId !== id) {
        this.currentTabId = id;
        this.$root.$emit('selectedTabChanged', id);
      }
    },
    createNewTabForUser(user) {
      this.tabs.push({id: user.id, title: user.name});
    }
  }
}
</script>

<style scoped>
ul {
  overflow-x: scroll;
}
</style>

