<template>
	<div>
		<el-carousel :interval="4000" type="card" height="650px">
			<el-carousel-item v-for="item in 6" :key="item">
				<h3 text="2xl" justify="center">{{ item }}</h3>
			</el-carousel-item>
		</el-carousel>
	</div>
</template>
  
<script>
export default {
	name: 'EmailShowTemplate'
}
</script>

<script setup>

</script>

<style>
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 700px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>