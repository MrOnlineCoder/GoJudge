<template>
	<b-alert variant="danger" :show="!!err">
			<font-awesome-icon icon="exclamation-circle"/>
			{{ err }}
	</b-alert>
</template>


<script >
export default {
	props: {
		err: {
			type: String
		}
	}
}	
</script>
