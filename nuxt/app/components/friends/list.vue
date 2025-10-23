<script setup>
const props = defineProps(["friends", "isLoading"]);
</script>

<template>
	<div v-if="!props.isLoading" class="columns-1 sm:columns-1 w-full">
		<friends-card
			v-for="friend in props.friends"
			:key="friend.id"
			:friend="friend"
		>
		</friends-card>
	</div>
	<!-- TODO поменять на прелоадер потому что скелетон хуета -->
	<div v-else class="columns-1 sm:columns-1 w-full">
		<friends-plug v-for="n in 2" :key="n"></friends-plug>
	</div>
</template>
