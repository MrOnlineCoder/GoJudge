<template>
	<div>
		<h1>Submissions</h1>
		<hr>
		<table class="table table-striped">
				<thead>
					<th>#</th>
					<th>User</th>
					<th>Problem</th>
					<th>Date</th>
					<th>Status</th>
				</thead>
				<tbody>
					<tr>
						<td>1</td>
						<td>admin</td>
						<td>A</td>
						<td>21.07.2019 22:19</td>
						<td class="text-danger">WRONG_ANSWER</td>
					</tr>
				</tbody>
		</table>
	</div>
</template>

<script>
export default {
	data() {
		return {
			list: [
				{

				}
			]
		}
	}
}
</script>