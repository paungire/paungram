<script setup>
import { ref, computed, onMounted } from "vue";

// активный таб
const activeTab = ref("online");

// имитация загрузки
const isLoading = ref(true);

// мок-друзья
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

const usernameHovered = ref(null);

const menuOpen = ref(false);

function toggleMenu(id) {
	setTimeout(() => {
		if (menuOpen.value === id) {
			menuOpen.value = false;
			return;
		}
		menuOpen.value = id;
	}, 100);
}

function closeMenu() {
	console.log("какого хуя");
	menuOpen.value = false;
}
</script>

<template>
	<UContainer class="py-8 max-w-3xl mx-auto">
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

			<!-- Скелетоны -->
			<div v-if="isLoading">
				<USkeleton
					v-for="n in 3"
					:key="n"
					class="h-16 w-full mb-3 rounded-xl"
				/>
			</div>

			<!-- Список друзей -->
			<div v-else class="flex flex-col gap-3">
				<div
					v-for="friend in filteredFriends"
					:key="friend.id"
					@mouseover="usernameHovered = friend.username"
					@mouseleave="usernameHovered = null"
					class="flex items-center justify-between p-3 hover:bg-muted max-h-16 rounded-xl cursor-pointer transition"
				>
					<div class="flex items-center gap-3">
						<UAvatar :src="friend.avatar" size="lg" />
						<div>
							<p class="font-medium">{{ friend.name }}</p>
							<p
								class="text-sm text-gray-500 transition hover:text-gray-300"
								:class="{
									'opacity-100': usernameHovered === friend.username,
									'opacity-0': usernameHovered !== friend.username,
								}"
							>
								{{ friend.username }}
							</p>
						</div>
					</div>
					<div class="flex items-center gap-2 relative">
						<btnSoft text="Сообщение" icon="mingcute:message-3-fill" />
						<btnSoft text="Позвонить" icon="mingcute:phone-fill" />
						<btnSoft
							text="Еще"
							icon="mingcute:more-2-fill"
							@click="toggleMenu(friend.id)"
						/>
						<div
							v-if="menuOpen === friend.id"
							class="menu absolute right-0 top-0 rounded-lg shadow-lg z-50 bg-[var(--ui-bg)] flex gap-1 p-2"
							v-click-outside="closeMenu"
						>
							<UButton
								variant="ghost"
								class="w-full justify-start"
								@click="closeMenu"
							>
								Пригласить
							</UButton>
							<UButton
								variant="ghost"
								color="error"
								class="w-full justify-start"
								@click="closeMenu"
							>
								Удалить
							</UButton>
						</div>
					</div>
				</div>
			</div>
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
	</UContainer>
</template>

<style scoped>
.hover\:opacity-100:hover {
	opacity: 1 !important;
}
</style>
