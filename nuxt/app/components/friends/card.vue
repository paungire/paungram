<script setup>
const props = defineProps(["friend"]);
const friend = props.friend;
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
	menuOpen.value = false;
}
</script>
<template>
	<div
		@mouseover="usernameHovered = friend.username"
		@mouseleave="usernameHovered = null"
		class="friend-card flex items-center justify-between p-3 hover:bg-muted max-h-16 rounded-2xl cursor-pointer transition break-inside-avoid bg-default not-last:mb-0.25 relative"
	>
		<div class="flex items-center gap-3">
			<UAvatar :src="friend.avatar" size="lg" />
			<div>
				<p class="font-medium">{{ friend.name }}</p>
				<p
					class="text-sm text-gray-300 transition hover:text-gray-500 dark:hover:text-gray-300 dark:text-gray-500"
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
		<div class="stick"></div>
	</div>
</template>

<style scoped>
.hover\:opacity-100:hover {
	opacity: 1 !important;
}
.stick {
	position: absolute;
	background-color: transparent;
	width: calc(100% - var(--spacing) * 6);
	height: 0.5px;
	bottom: -1px;
	left: calc(var(--spacing) * 3);
	transition: background-color 0.2s ease;
}
.friend-card:not(:last-child):not(:hover):not(:has(+ .friend-card:hover))
	.stick {
	background-color: var(--border);
}
</style>
