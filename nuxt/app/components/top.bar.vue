<script setup>
const route = useRoute();
</script>

<template>
	<div class="relative">
		<div class="flex items-center text-center justify-center">
			{{ route.meta.title ?? "Заголовок" }}
		</div>
		<stick :passive="true" />
	</div>
</template>
