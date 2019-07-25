<template>
	<div>
		<h1>Scoreboard</h1>
		<hr>
		<table class="table table-striped">
			<thead>
				<th>Rank</th>
				<th>User</th>
				<th>A</th>
				<th>B</th>
				<th>C</th>
				<th>Total</th>
			</thead>
			<tbody>
				<tr>
					<td>1</td>
					<td>user01</td>
					<td>
						<font-awesome-icon icon="star" class="text-success"/>
						<b>100</b>
					</td>
					<td>
						<font-awesome-icon icon="star-half-alt" class="text-warning"/>
						56
					</td>
					<td>
						<font-awesome-icon icon="times" class="text-danger"/>
						0
					</td>
					<td>
						156
					</td>
				</tr>
			</tbody>
		</table>	
	</div>
</template>

<script>
export default {

}
</script>