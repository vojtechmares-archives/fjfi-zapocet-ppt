<template>
  <div>
    <div class="mt-3">
      <transition name="fade">
        <div v-if="!hidden" id="messages" class="rounded">
          <div v-for="message in messages" class="flex" :class="{ 'justify-end': message.sentByMe, 'justify-start': !message.sentByMe}">
            <p class="m-1 max-w-md rounded px-3 py-1"
              :class="{ 'bg-gray-200': message.sentByMe && !message.system,  'bg-blue-200': !message.sentByMe && !message.system, 'bg-red-600': message.system }"
              >
                {{ message.text }}
              </p>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedTabId: null,
      hidden: false,
      messages: [
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        {sentByMe: true, text: 'Hello John!'},
        {sentByMe: false, text: 'How you been? I haven\'t heard about you for a while, is everything ok?'},
        {sentByMe: false, text: 'Sorry for me asking, but it has been 9 months!'},
        {sentByMe: false, text: 'Evan?'},
        //{sentByMe: false, system: true, text: 'Error: lost connection'}
      ]
    }
  },
  mounted() {
    this.fetchMessages();
    this.$root.$on('selectedTabChanged', userId => {
      this.selectedTabId = userId;
      console.log('selected tab id ' + this.selectedTabId);
      this.hidden = true;
      setTimeout(() => {
        this.hidden = false;
      }, 500);
    });
    this.$root.$on('ws.data.recieved', data => {
      if (data.from === this.sender.id) {}
    });

    this.scrollDown();
  },
  updated() {
    this.scrollDown();
  },
  methods: {
    fetchMessages(userId) {
      // GET /api/v1/messages?userId=<userId>&latest=20
    },
    scrollDown() {
      const el = this.$el.querySelector('#messages');
      el.scrollTop = el.scrollHeight;
    }
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity .25s;
}

.fade-enter, .fade-leave-to {
  opacity: 0;
}

#messages {
  max-height: 33rem;
  overflow-y: auto;
}
</style>
