<template>
	<div>
		<div class="add-container">
			<b-button variant="success" size="lg" @click="openNewProblemDialog()">
				<font-awesome-icon icon="plus"/>
				Create Problem
			</b-button>
		</div>
		<hr>
		<ErrorBlock :err="error"/>
		<table class="table">
			<thead>
				<th>ID</th>
				<th>Name</th>
				<th>Time limit</th>
				<th>Memory limit</th>
				<th>Edit</th>
				<th>Delete</th>
			</thead>
			<tbody>
				<tr v-for="problem in problems" :key="problem.id">
					<td>{{problem.id}}</td>
					<td>{{problem.name}}</td>
					<td>{{problem.timelimit}} ms</td>
					<td>{{problem.memlimit}} KB</td>
					<td>
						<b-button variant="warning" @click="openEditProblemDialog(problem)">
							<font-awesome-icon icon="pencil-alt"/>
						</b-button>
					</td>
					<td>
						<b-button variant="danger" @click="openDeleteProblemDialog(problem.id)">
							<font-awesome-icon icon="trash"/>
						</b-button>
					</td>
				</tr>
			</tbody>
		</table>

		<b-modal
			:header-bg-variant="isEditing ? 'warning' : 'success'"
			title="Create/edit problem"
			ref="createProblemModal"
			ok-title="Cancel"
			ok-only
			no-close-on-backdrop
			size="lg">
			<ErrorBlock :err="error"/>
			<b-form @submit.prevent="submitProblemData">
	      <b-form-group label="Name:">
	        <b-form-input
	          v-model="problemData.name"
	          required
	          placeholder="Enter name for problem"
	        ></b-form-input>
	      </b-form-group>

	      <b-form-group label="Time limit:">
	      	<b-input-group append="ms." type="number">
				    <b-form-input v-model="problemData.timelimit" required/>
				  </b-input-group>
				  <br>
				  Use preset value for time limit:
				  <b-button-group size="sm">
				    <b-button @click="setTimelimitPreset(100)">100 ms</b-button>
				    <b-button @click="setTimelimitPreset(250)">250 ms</b-button>
				    <b-button @click="setTimelimitPreset(500)">500 ms</b-button>
				    <b-button @click="setTimelimitPreset(1000)">1 second</b-button>
				    <b-button @click="setTimelimitPreset(2000)">2 seconds</b-button>
				    <b-button @click="setTimelimitPreset(3000)">3 seconds</b-button>
				    <b-button @click="setTimelimitPreset(5000)">5 seconds</b-button>
				  </b-button-group>
	      </b-form-group>

	      <b-form-group label="Memory limit:">
	      	<b-input-group append="KB" type="number">
				    <b-form-input v-model="problemData.memlimit" required/>
				  </b-input-group>
				  <br>
				  Use preset value for memory limit:
				  <b-button-group size="sm">
				    <b-button @click="setMemlimitPreset(1)">1 MB</b-button>
				    <b-button @click="setMemlimitPreset(16)">16 MB</b-button>
				    <b-button @click="setMemlimitPreset(64)">64 MB</b-button>
				    <b-button @click="setMemlimitPreset(128)">128 MB</b-button>
				    <b-button @click="setMemlimitPreset(256)">256 MB</b-button>
				    <b-button @click="setMemlimitPreset(512)">512 MB</b-button>
				    <b-button @click="setMemlimitPreset(1024)">1 GB</b-button>
				  </b-button-group>
	      </b-form-group>

	      <b-tabs>
	      	<b-tab title="Problem text" active>
	      		<br>
	      		<b-form-group>
			      	 <b-form-textarea
						      v-model="problemData.text"
						      placeholder="Enter statment/text for the problem."
						      rows="5"
						    ></b-form-textarea>
						    <p>
						    	<b>Note: </b> <a href="https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet">Markdown</a> is supported in problem text.
						    </p>
			      </b-form-group>
	      	</b-tab>
	      	<b-tab title="Preview">
	      		<div class="problem-preview">
	      			<div v-html="renderedProblemText"></div>
	      		</div>
	      	</b-tab>	
	      </b-tabs>

	      <b-button variant="success" type="submit">
	      	<font-awesome-icon icon="save"/>
	      	{{ isEditing ? 'Save' : 'Create'}}
	      </b-button>
	    </b-form>
		</b-modal>

		<b-modal
			header-bg-variant="danger"
			title="Delete problem"
			ok-only
			no-close-on-backdrop
			ok-title="Cancel"
			ref="deleteProblemModal">
			<ErrorBlock :err="error"/>
			<p>Are you sure you want to delete this problem?</p>
			<p><b>All queued submissions for this problem will be forcefully cancelled.</b></p>
			<b>This action is irreversible.</b>
			<hr>
			<b-button variant="danger" @click="deleteProblem()">
				<font-awesome-icon icon="trash"/>
				Yes, delete this problem.
			</b-button>
		</b-modal>
	</div>
</template>

<script>
import marked from 'marked'
import axios from 'axios'

export default {
	data() {
		return {
			problems: [],
			problemData: {
				id: -1,
				name: null,
				timelimit: 0,
				memlimit: 0,
				text: null
			},
			deleteData: {
				id: -1
			},
			isEditing: false,
			error: null
		}
	},
	methods: {
		openNewProblemDialog() {
			this.$refs.createProblemModal.show();
			this.problemData = {
				id: -1,
				name: null,
				timelimit: this.problemData.timelimit,
				memlimit: this.problemData.memlimit,
				text: null
			};
			this.isEditing = false;
		},
		openEditProblemDialog(problem) {
			this.problemData = problem;
			this.isEditing = true;
			this.$refs.createProblemModal.show();
		},
		openDeleteProblemDialog(id) {
			this.deleteData.id = id;
			this.$refs.deleteProblemModal.show();
		},
		setMemlimitPreset(megs) {
			this.problemData.memlimit = megs * 1024;
		},
		setTimelimitPreset(ms) {
			this.problemData.timelimit = ms;
		},
		submitProblemData() {
			//Little hack to minimize code
			//Both endpoints receive same body input
			if (this.isEditing) {
				this.saveProblem('/api/admin/editProblem');
			} else {
				this.saveProblem('/api/admin/createProblem');
			}
		},
		saveProblem(endpoint) {
			this.validateProblemNumbers();

			axios.post(endpoint, {
				problem: this.problemData
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.fetchProblems();

				this.$refs.createProblemModal.hide();

				this.error = null;
			}).catch(error => {
				this.error = `Save Problem Request failed: ${error}`
			});
		},
		deleteProblem() {
			this.error = null;

			axios.post('/api/admin/deleteProblem', {
				problem_id: this.deleteData.id
			}).then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.fetchProblems();

				this.$refs.deleteProblemModal.hide();

				this.error = null;
			}).catch(error => {
				this.error = `Delete User Request failed: ${error}`
			});
		},
		fetchProblems() {
			axios.get('/api/admin/problems').then(response => {
				if (!response.data.success) {
					this.error = response.data.message;
					return;
				}

				this.problems = response.data.problems;
			});
		},
		validateProblemNumbers() {
			this.problemData.timelimit = parseInt(this.problemData.timelimit);
			this.problemData.memlimit = parseInt(this.problemData.memlimit);
		}
	},
	computed: {
		renderedProblemText() {
			if (!this.problemData.text) {
				return '';
			}

			return marked(this.problemData.text);
		}
	},
	mounted() {
		this.fetchProblems();
	}
}
</script>

<style scoped>
.add-container {
	text-align: center;
}

.problem-preview {
	min-height: 120px;
	width: 100%;
}
</style>