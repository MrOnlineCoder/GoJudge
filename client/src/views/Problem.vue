<template>
	<div>
		<div class="problem-header">
			<h2>{{ problem.name }}</h2>
			<small>Time Limit: <b>{{ problem.timelimit | formatTimelimit}}</b> </small>
			<br>
			<small>Memory Limit: <b>{{ problem.memlimit | formatMemlimit}}</b> </small>
			<hr>
		</div>
		<div class="problem-text">
			<div v-html="renderedProblemText"></div>
		</div>
		<hr>
	</div>
</template>

<script>
import marked from 'marked'

export default {
	data() {
		return {
			problem: {
				name: "",
				timelimit: "",
				memlimit: "",
				text: ""
			}
		}
	},
	methods: {
		fetchProblem() {
			this.problem = {
				name: "B - Оливер и минимум",
				timelimit: 2000,
				memlimit: 131072,
				text: `Оливер поспорил со своим другом Джеком, что сможет сделать массив с самым минимальным числом. Джек не верит Оливеру, и требует программы, которая найдет минимум в массиве и подтвердит умение Оливера. 

На вход дается число **n**, (1 < **n** < 10000). Дальше идут **n** чисел. Выведите минимальное число из этих **n** чисел. 

Пример:

Ввод:
\`\`\`
4
8 3 9 5
\`\`\`

Вывод:
\`\`\`
3
\`\`\``
			}
		}
	},
	computed: {
		renderedProblemText() {
			return this.problem.text ? marked(this.problem.text) : '';
		}
	},
	mounted() {
		this.fetchProblem();
	}
}
</script>

<style scoped>
.problem-header {
	text-align: center;
}
</style>