<template>
	<div>
		<h1>Admin</h1>
		<hr>
		<b-tabs content-class="mt-3">
	    <b-tab title="Configuration" active>
	    	<Configuration />
	    </b-tab>
	    <b-tab title="Contest">
	    	<ContestEditor />
	    </b-tab>
	    <b-tab title="Users">
	    	<UsersEditor />
	    </b-tab>
	    <b-tab title="Problems">
	    	<ProblemsEditor />
	    </b-tab>
	    <b-tab title="Tests">
	    	<TestsEditor />
	    </b-tab>
  	</b-tabs>
	</div>
</template>

<script>
import UsersEditor from '@/components/admin/UsersEditor.vue'
import ProblemsEditor from '@/components/admin/ProblemsEditor.vue'
import TestsEditor from '@/components/admin/TestsEditor.vue'
import ContestEditor from '@/components/admin/ContestEditor.vue'
import Configuration from '@/components/admin/Configuration.vue'

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
			TestsEditor,
			ProblemsEditor,
			ContestEditor,
			Configuration
		}
	}
</script>