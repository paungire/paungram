<script setup>
const toast = useToast();

function showSuccess() {
	toast.add({
		title: "Успешно!",
		description: "Данные сохранены",
		color: "success",
	});
}

function showError() {
	toast.add({
		title: "Ошибка",
		description: "Не удалось сохранить данные",
		color: "error",
	});
}
</script>

<template>
	<div>
		<UButton color="success" @click="showSuccess">Показать успех</UButton>
		<UButton color="error" @click="showError">Показать ошибку</UButton>
	</div>
</template>
