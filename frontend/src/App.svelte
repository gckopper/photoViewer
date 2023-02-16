<script lang="ts">
    import { onMount } from 'svelte';
    import {Next, Priv, First} from '../wailsjs/go/main/App.js'
    import {WindowSetTitle} from '../wailsjs/runtime/runtime.js'

  let resultText: string = "loading.png"
  let format: string = "png"

  function processShit(image:Array<string>) {
    format = image[0];
    resultText = image[1];
    console.log(resultText);
    WindowSetTitle(image[2]);
  }

  function next(): void {
    Next().then(result => {
      processShit(result)
    })
  }
  function priv(): void {
    Priv().then(result => {
      processShit(result)
    })
  }
  function first(): void {
    First().then(result => {
      processShit(result)
    })
  }
  function change(): void {
    let element = document.getElementById("logo");
    if (element.style.maxHeight == "fit-content") {
      element.style.maxHeight = "100vh"
    } else {
      element.style.maxHeight = "fit-content";
    }
  }
  function resetScroll(): void {
    window.scrollTo(window.scrollX, 0)
  }
  window.addEventListener("keydown", function (event) {
    switch (event.key) {
      case "ArrowLeft":
        priv()
        event.preventDefault();
        return;
      case "ArrowRight":
        next()
        event.preventDefault();
        return;
      case " ":
        change()
        event.preventDefault();
        return;
      case "ArrowDown":
      case "ArrowUp":
        break;
      default:
        return;
    }
    if (window.innerHeight + Math.ceil(window.pageYOffset) >= document.body.scrollHeight) { 
      if (event.key == "ArrowDown") {
        next()
        event.preventDefault();
      }
    }
    if (Math.ceil(window.pageYOffset) <= 0) {
      if (event.key == "ArrowUp") {
        priv()
        event.preventDefault();
      }
    }
}, true);

onMount(() => {
  first()
})
</script>

<main>  
  <div class="input-box big" id="input">
    <button class="btn-left big no-border-bs" on:click={priv}>
      &lt;
    </button>
    <button class="btn-main big no-border-bs" on:click={change}/>
    <button class="btn-right big no-border-bs" on:click={next}>
      &gt;
    </button>
  </div>
  <img 
    class="no-border-bs"
    alt="User provided content" on:load={resetScroll} 
    id="logo" 
    src="data:{format};base64,{resultText}"
  >
</main>

<style>
  .big {
    height: 100%;
    width: 100%;
  }
  .no-border-bs {
    outline: 0;
    border: none;
    padding: 0;
    margin: 0;
  }

  main {
    display: flex;
    justify-content: center;
    height: 100vh;
  }

  #logo {
    display: block;
    max-width:100%;
    max-height: 100vh;
    margin: auto;
  }

  .input-box {
    position: fixed;
    display: flex;
  }

  .input-box * {
    background-color: transparent;
    cursor: pointer;
    font-size: 12rem;
    color: transparent;
  }

  .input-box .btn-right:hover, .input-box .btn-left:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
    opacity: 30%;
    display: inline;
  }

</style>
