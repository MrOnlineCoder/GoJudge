<template>
	<div>
		<ErrorBlock :err="error"/>
		<b-form-group>
			<b-form-select v-model="selectedProblem">
				<option :value="null" disabled>--- Choose a problem to edit tests for ---</option>
				<option v-for="p in problems" :key="p.id" :value="p.id">{{p.name}}</option>
			</b-form-select>			
		</b-form-group>
		<div v-if="selectedProblem === null">
			<p>
				Select a problem from the list above to edit tests for it.
			</p>
		</div>
		<b-card v-if="selectedProblem !== null">
			<div class="add-container">
				<b-button variant="success" size="lg" @click="openNewTestDialog()">
				<font-awesome-icon icon="plus"/>
				Add Test
			</b-button>
			</div>
			<hr>
			<table class="table" v-if="tests.length > 0">
				<thead>
					<th>ID</th>
					<th>Index</th>
					<th>Checking Method</th>
					<th>Is Sample</th>
					<th>Input/Output</th>
					<th>Edit</th>
					<th>Delete</th>
				</thead>
				<tbody>
					<tr v-for="test in tests" :key="test.id">
						<td>{{test.id}}</td>
						<td>{{test.test_index}}</td>
						<td>{{check_methods[test.check_method]}}</td>
						<td>{{test.is_sample ? 'Yes' : 'No'}}</td>
						<td>{{test.input.length}} B / {{test.output.length}} B</td>
						<td>
							<b-button variant="warning" @click="openEditTestDialog(test)">
								<font-awesome-icon icon="pencil-alt"/>
							</b-button>
						</td>
						<td>
							<b-button variant="danger" @click="openDeleteTestDialog(test.id)">
								<font-awesome-icon icon="trash"/>
							</b-button>
						</td>
					</tr>
				</tbody>
			</table>
		</b-card>

		<b-modal
			header-bg-variant="success"
			title="Create/edit new test"
			ref="createTestModal"
			ok-title="Cancel"
			ok-only
			no-close-on-backdrop
			size="lg">
			<ErrorBlock :err="error"/>
			<b-form @submit.prevent="submitTest">
	      <b-form-group label="Problem ID:">
	        <b-form-input
	          v-model="testData.problem_id"
	          required
	          disabled
	        ></b-form-input>
	      </b-form-group>

	      <b-form-group label="Index:" 
	      	description="Index determines order of tests during the submission check. You may leave it without changing">
	        <b-form-input
	          v-model="testData.test_index"
	          required
	        ></b-form-input>
	      </b-form-group>

	      <b-form-group label="Checking method:">
	        <b-form-radio-group v-model="testData.check_method" stacked>
	        	<b-form-radio v-for="v,k in check_methods" :key="k" :value="k">{{v}}</b-form-radio>
	        </b-form-radio-group>
	        <kbd>Checking method description:</kbd>
	        <p v-if="testData.check_method === 0">
	        	Jury's output and participant's solution output will be compared byte-by-byte, character-by-character. This method can be used when there is only one (usually short) single solution to a problem. Please note, that despite of being called strict, whitespace and newline characters at the end of the output are trimmed.
	        </p>
	        <p v-if="testData.check_method === 1">
	        	Jury's output and participant's solution output will be split into <b>tokens</b> and the outputs will be compared token-by-token. This method of checking can be used when there is also only one solution to the problem, but can have complicated output. This way of checking ignores any whitespace or newline characters, allowing participants to present their answer in different ways.
	        </p>
	        <p v-if="testData.check_method === 2">
	        	External program called <b>checker</b> is called for checking participant's answer. Use this method of checking for complex solutions, with multiple answers and complex output. 
	  				Please refer documentation for info on checkers.
	        </p>
	      </b-form-group>

	      <b-form-group description="If this test is sample, it will be displayed publicly in problem text.">
	        <b-form-checkbox v-model="testData.is_sample">
	        	Is sample test
	        </b-form-checkbox>
	      </b-form-group>

	      <b-form-group label="Input">
	      	 <b-form-textarea
						  v-model="testData.input"
						  placeholder="Enter input data for this test."
						  rows="5"
						></b-form-textarea>
	      </b-form-group>

	      <b-form-group label="Output" v-if="testData.check_method !== 2">
	      	 <b-form-textarea
						  v-model="testData.output"
						  placeholder="Enter expected output data for this test."
						  rows="5"
						></b-form-textarea>
	      </b-form-group>

	      <b-button variant="success" type="submit">
	      	<font-awesome-icon icon="save"/>
	      	{{ isEditing ? 'Save': 'Create'}}
	      </b-button>
	    </b-form>
		</b-modal>
	</div>
</template>

<script>
import axios from 'axios'

export default {
	data() {
		return {
			problems: [],
			tests: [],
			selectedProblem: null,
			error: null,
			testData: {
				problem_id: -1,
				test_index: 0,
				check_method: 0,
				checker_id: 0,
				is_sample: false,
				input: null,
				output: null
			},
			isEditing: false,
			check_methods: {
				0: 'Strict comparison',
				1: 'Token comparison',
				2: 'Checker program'
			}
		}
	},
	methods: {
		fetchProblems() {
			axios.get('/api/admin/problems').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.problems = response.data.problems;
			});
		},
		fetchTests() {
			if (this.selectedProblem === null) return; 

			this.error = null;

			axios.get('/api/admin/tests/'+this.selectedProblem).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.tests = response.data.tests;
			});
		},
		openNewTestDialog() {
			this.testData.problem_id = this.selectedProblem;
			this.testData.test_index = this.tests.length+1;
			this.isEditing = false;
			this.$refs.createTestModal.show();
		},
		openEditTestDialog(test) {
			this.testData = test;
			this.isEditing = true;
			this.$refs.createTestModal.show();
		},

		saveTest(endpoint) {
			axios.post(endpoint, {
				test: this.testData
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.fetchTests();

				this.$refs.createTestModal.hide();

				this.error = null;
			}).catch(error => {
				this.error = `Create/Save Test Request failed: ${error}`
			});
		},
		submitTest() {
			if (this.isEditing) {
				this.saveTest('/api/admin/tests/editTest');
			} else {
				this.saveTest('/api/admin/tests/createTest');
			}
		}
	},
	watch: {
		selectedProblem() {
			this.fetchTests();
		}
	},
	mounted() {
		this.fetchProblems();
	}
}
</script>

<style>
.add-container {
	text-align: center;
}
</style>