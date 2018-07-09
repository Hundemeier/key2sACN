<template>
  <div class="mx-auto text-center">
    {{message}}
    <div v-if="error" class="text-danger">Error on websocket!</div>
    <div v-if="closed" class="text-danger">WebSocket closed!</div>
  </div>
</template>

<script>
export default {
  name: 'KeyEvent',
  data () {
    return {
      error: false,
      closed: true,
      message: ''
    }
  },
  mounted: function () {
    // const addr = 'ws://192.168.1.20:8080/websocket'
    const addr = 'ws://' + window.location.host + '/websocket'
    console.log('start websocket')
    var ws = new WebSocket(addr)

    ws.onerror = function () {
      console.log('Error on websocket')
      this.error = true
    }.bind(this)

    ws.onopen = function () {
      console.log('websocket was opened!')
      this.error = false
      this.closed = false
      this.message = 'WebSocket opened!'
    }.bind(this)

    ws.onclose = function () {
      this.closed = true
      console.log('websocket was closed! Try to reconnect...')
      ws = new WebSocket(addr)
      // eslint-disable-next-line
    }.bind(this)

    ws.onmessage = function (e) {
      var event = JSON.parse(e.data)
      if (event.Type === 1) {
        if (event.Data.Value === 0) {
          this.message = event.Data.KeyboardID + ':' + event.Data.KeyCode + ' Key Up'
        } else if (event.Data.Value === 1) {
          this.message = event.Data.KeyboardID + ':' + event.Data.KeyCode + ' Key Down'
        }
      }
    }.bind(this)
  }
}
</script>
