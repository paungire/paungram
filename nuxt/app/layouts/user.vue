<template>
	<div>
		<top-bar />
		<div class="flex gap-2">
			<menu-left />
			<div class="w-full flex flex-col">
				<slot />
			</div>
		</div>
	</div>
</template>
