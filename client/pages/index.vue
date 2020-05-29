<template>
  <div>
    <h1 class="container mt-1">
      Chat
    </h1>
    <div id="chat" class="container bg-white shadow-md rounded flex flex-col">
      <div class="grid grid-cols-12 gap-4">
        <div class="col-span-9">
          <TabSelector />
          <Messages />
        </div>
        <div class="col-span-3">
          <OnlineUsers />
        </div>
      </div>
      <div class="grid grid-cols-12 gap-4 mt-auto">
        <div class="col-span-9">
          <div class="inline-flex w-full mt-3">
            <input type="text" name="chat" class="shadow appearance-none border rounded-l w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline">
            <button class="shadow bg-blue-500 hover:bg-blue-400 text-white font-bold py-2 px-4 hover:border-blue-500 rounded-r">Send</button>
          </div>
        </div>
        <div class="col-span-3">
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TabSelector from '~/components/TabSelector.vue';
import OnlineUsers from '~/components/OnlineUsers.vue';
import Messages from '~/components/Messages.vue';

export default {
  components: {
    Messages: Messages,
    TabSelector: TabSelector,
    OnlineUsers: OnlineUsers,
  },
  data() {
    return {
      ws: null,
      user: {
        id: null,
        name: '',
      }
    }
  },
  created() {
    this.ws = new WebSocket('ws://localhost:3000/chat/ws');

    this.ws.onopen = function(event) {
      console.log('Successfully connected to WebSocket server!');
    };

    this.ws.onmessage = function(event) {
      console.log('[DEBUG] Recieved data: ' + event);
      this.$root.$emit('ws.data.recieved', event.data);
    }

    this.ws.onerror = function(event) {
      console.log('WS error: ' + error);
    }
  },
  mounted() {
    this.$root.$on()
  },
  methods: {
    sendMessage(toUserId, messageText) {
      this.ws.send(JSON.stringify({receiverId: toUserId, message: messageText}));
    },
    sendCommand(commandText) {
      let commandSplit = commandText.split(' ');
      commandText = commandSplit[0];

      let argsObj = {};
      for (let i=1; i < commandSplit.length; i += 2) {
        const keys = [...commandSplit.keys()];

        if (keys.includes(i) && keys.includes(i+1)) {
          argsObj.commandSplit[i] = commandSplit[i+1];
        }
      }

      this.ws.send(JSON.stringify({command: commandText, args: argsObj}))
    },
    userSelected(user) {
      this.components.TabSelector.$emit('createNewTabForUser', user);
    },
  }
}
</script>

<style>
body {
  background-color: rgb(237, 242, 247);
}

#chat {
  padding: 1rem;
  height: 85vh;
}

/* Sample `apply` at-rules with Tailwind CSS
.container {
  @apply min-h-screen flex justify-center items-center text-center mx-auto;
}
*/
.container {
  margin: 0 auto;
  width: 60rem;
}

h1 {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  letter-spacing: 1px;
  font-size: 3rem;
}
</style>
