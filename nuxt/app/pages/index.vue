<script>
definePageMeta({
	middleware: "welcome",
});
</script>
<template>
	<div>
		<main
			class="w-full max-w-[var(--ui-container)] p-4 sm:p-6 lg:p-8 h-full items-center mx-auto"
		>
			<div class="columns-1 sm:columns-1 lg:columns-2 [column-gap:2rem] w-full">
				<dashboard-card link="/id/1" title="Профиль">
					<p>Вы онлайн 3 часа сегодня</p>
				</dashboard-card>

				<dashboard-card link="/id/1/friends" title="Друзья">
					<p>Онлайн: 5 / Всего: 20</p>
				</dashboard-card>

				<dashboard-card link="/id/1/friends" title="Сообщения">
					<p>Непрочитанных: 12</p>
					<p>Непрочитанных: 12</p>
					<p>Непрочитанных: 12</p>
					<p>Непрочитанных: 12</p>
				</dashboard-card>

				<dashboard-card link="/id/1/friends" title="Сервера">
					<p>Вы состоите в 3 серверах</p>
					<p>Вы состоите в 3 серверах</p>
					<p>Вы состоите в 3 серверах</p>
					<p>Вы состоите в 3 серверах</p>
				</dashboard-card>
			</div>
		</main>
	</div>
</template>
