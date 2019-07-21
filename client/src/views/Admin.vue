<template>
	<div>
		<h1>Admin</h1>
		<hr>
		<b-tabs content-class="mt-3">
	    <b-tab title="General" active>
	    	
	    </b-tab>
	    <b-tab title="Status">
	    	
	    </b-tab>
	    <b-tab title="Users">
	    	<UsersEditor />
	    </b-tab>
	    <b-tab title="Problems">
	    	<ProblemsEditor />
	    </b-tab>
  	</b-tabs>
	</div>
</template>

<script>
import UsersEditor from '@/components/admin/UsersEditor.vue'
import ProblemsEditor from '@/components/admin/ProblemsEditor.vue'

	export default {
		data() {
			return {

			}
		},
		methods: {

		},
		mounted() {
			if (!this.$session.isLoggedIn() || this.$session.getCachedUser().access < 2) {
				this.$router.push('/');
			}
		},
		components: {
			UsersEditor,
			ProblemsEditor
		}
	}
</script>