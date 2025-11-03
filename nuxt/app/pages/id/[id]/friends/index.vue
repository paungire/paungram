<script setup>
definePageMeta({
	layout: "user",
	title: "Друзья",
});

// активный таб
const route = useRoute();
const url = useRequestURL();

const tab = url.searchParams.get("tab");
const activeTab = ref(tab || "online");

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

function changeTab(tab) {
	activeTab.value = tab;
	let url = new URL(window.location.href);
	url.searchParams.set("tab", tab);
	window.history.replaceState({}, "", url);
}
</script>

<template>
	<div class="flex flex-col gap-0.25 pr-2">
		<!-- Tabs -->
		<div class="flex items-center gap-3 py-4 relative">
			<button
				:variant="activeTab === 'online' ? 'soft' : 'ghost'"
				color="neutral"
				@click="changeTab('online')"
			>
				В сети
			</button>
			<button
				:variant="activeTab === 'all' ? 'soft' : 'ghost'"
				color="neutral"
				@click="changeTab('all')"
			>
				Все
			</button>
			<button
				:variant="activeTab === 'add' ? 'soft' : 'solid'"
				color="primary"
				icon="i-lucide-user-plus"
				@click="changeTab('add')"
			>
				Добавить в друзья
			</button>
			<stick passive="true" />
		</div>

		<!-- Search (только для первых двух табов) -->
		<div v-if="activeTab !== 'add'">
			<input
				placeholder="Поиск друзей..."
				icon="i-lucide-search"
				size="md"
				class="py-4 w-full mb-0.25"
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
			<input
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
