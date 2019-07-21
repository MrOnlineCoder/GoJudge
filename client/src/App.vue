<template>
  <div class="gojudge-app container">
    <TheNavbar />
    <router-view/>
  </div>
</template>

<script>
  import TheNavbar from '@/components/TheNavbar'

  export default {
    components: {
      TheNavbar
    }
  }
</script>
