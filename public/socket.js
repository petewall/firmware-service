/* global io */

const socket = io()
socket.on('connect', () => {
  console.log('Connected to Server')  
})
