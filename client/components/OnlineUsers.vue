<template>
  <div>
    <h2 class="py-2 font-semibold">Online users</h2>
    <p class="text-xs text-gray-400">
      Click to open new tab
    </p>
    <ul>
      <li v-for="user in users">
        <button class="w-full bg-gray-300 px-2 py-2 mt-2 hover:bg-gray-400 text-gray-800 font-bold rounded"
          @click="selectUser(user)">
          {{ user.name }}
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  mounted() {
    this.fetchOnlineUsers(); // fetch online users right away so it is up to date
    this.interval = setInterval(() => {
      this.fetchOnlineUsers();
      }, 5000); // every 5 seconds, get online user list
  },
  methods: {
    fetchOnlineUsers() {
      //console.log('[DEBUG] Fetching users online');
      // GET /api/v1/users/online
    },
    selectUser(user) {
      this.$root.$emit('selectUser', user);
    }
  },
  data() {
    return {
      interval: null,
      users: [
        {id: 'john-doe', name:'John Doe'},
        {id: 'emma-whatson', name: 'Emma Watson'},
        {id: 'erica-albright', name:'Erica Albright'},
        {id: 'ohhimark', name:'Mark Zuckerberg'},
        {id: 'bradpitt', name:'Brad Pitt'},
        {id: 'angelina', name:'Angelina Jolie'},
        {id: 'mamut', name:'Mamut'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        {id: 'john-doe', name:'John Doe'},
        ]
    }
  }
}
</script>

<style scoped>
ul {
  max-height: 33rem;
  overflow-y: auto;
}
</style>
