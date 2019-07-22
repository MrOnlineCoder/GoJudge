<template>
  <div class="gojudge-app container">
    <TheNavbar />
    <router-view/>
  </div>
</template>

<script>
import TheNavbar from '@/components/TheNavbar'
import ContestService from '@/services/contest'

export default {
  components: {
    TheNavbar
  },
  created() {
    ContestService.loadStatus();
  }
}
</script>
