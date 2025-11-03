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
	<div v-else>
		<Icon name="mingcute:loading-fill" size="80" class="m-auto spin"></Icon>
	</div>
</template>
