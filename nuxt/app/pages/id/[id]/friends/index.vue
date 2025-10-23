<script setup>
definePageMeta({
	layout: "user", // имя layout-файла
});

// активный таб
const activeTab = ref("online");

// имитация загрузки
const isLoading = ref(true);

// мои-друзья
const friends = ref([]);

// фильтруем по табу
const filteredFriends = computed(() => {
	if (activeTab.value === "online")
		return friends.value.filter((f) => f.online);
	return friends.value;
});

// при маунте имитируем API
onMounted(() => {
	setTimeout(() => {
		friends.value = [
			{
				id: 1,
				name: "Алексей Морозов",
				username: "@frosty",
				avatar: "https://i.pravatar.cc/80?img=1",
				online: true,
			},
			{
				id: 2,
				name: "Мария Иванова",
				username: "@maria",
				avatar: "https://i.pravatar.cc/80?img=2",
				online: false,
			},
			{
				id: 3,
				name: "Кирилл Лебедев",
				username: "@lebedev",
				avatar: "https://i.pravatar.cc/80?img=3",
				online: true,
			},
		];
		isLoading.value = false;
	}, 1500);
});
</script>

<template>
	<div>
		<!-- Tabs -->
		<div class="flex items-center gap-3 mb-6">
			<UButton
				:variant="activeTab === 'online' ? 'soft' : 'ghost'"
				color="neutral"
				@click="activeTab = 'online'"
			>
				В сети
			</UButton>
			<UButton
				:variant="activeTab === 'all' ? 'soft' : 'ghost'"
				color="neutral"
				@click="activeTab = 'all'"
			>
				Все
			</UButton>
			<UButton
				:variant="activeTab === 'add' ? 'soft' : 'solid'"
				color="primary"
				icon="i-lucide-user-plus"
				@click="activeTab = 'add'"
			>
				Добавить в друзья
			</UButton>
		</div>

		<!-- Search (только для первых двух табов) -->
		<div v-if="activeTab !== 'add'">
			<UInput
				placeholder="Поиск друзей..."
				icon="i-lucide-search"
				size="lg"
				class="mb-6 w-full"
			/>

			<!-- Список друзей -->
			<friends-list
				:friends="filteredFriends"
				:isLoading="isLoading"
			></friends-list>
		</div>

		<!-- Добавить друга -->
		<div v-else>
			<h2 class="text-xl font-bold mb-4">Добавить друга</h2>
			<UInput
				placeholder="Введите никнейм или ID..."
				icon="i-lucide-user-search"
				size="lg"
				class="w-full"
			/>
			<p class="text-sm text-gray-500 mt-2">
				Отправь запрос пользователю, чтобы добавить его в друзья.
			</p>
		</div>
	</div>
</template>
