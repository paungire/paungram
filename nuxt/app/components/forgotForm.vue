<script>
import * as z from "zod";
z.config(z.locales.ru());

export default {
	data() {
		return {
			error: null,

			schema: z.object({
				email: z.string().email(),
			}),

			state: {
				email: "",
			},
		};
	},
	methods: {
		onSubmit(payload) {
			// TODO
			console.log("Submitted", payload);
			if (this.error) {
				this.error = null;
			} else {
				this.error = "Произошла ошибка, напишите на почту paungire@bk.ru";
			}
		},
	},
};
</script>

<template>
	<div class="p-4">
		<UPageCard class="w-full max-w-md">
			<div>
				<div class="flex flex-col text-center items-center mb-6">
					<div class="mb-2">
						<UIcon name="pepicons-print:raise-hand" class="size-12"></UIcon>
					</div>
					<div class="text-xl text-pretty font-semibold text-highlighted">
						Восстановление пароля
					</div>
					<div class="mt-1 text-base text-pretty text-muted">
						Вспомнили пароль?
						<ULink to="/auth/" class="text-primary font-medium"
							>Повторить вход</ULink
						>.
					</div>
				</div>

				<UForm
					:schema="schema"
					:state="state"
					class="space-y-4"
					@submit="onSubmit"
				>
					<UFormField
						name="email"
						label="Email указанный при регистрации"
						required
					>
						<UInput
							v-model="state.email"
							placeholder="Введите email"
							class="w-full"
						/>
					</UFormField>

					<UAlert
						v-if="error"
						color="error"
						icon="i-lucide-info"
						:title="error"
					/>

					<UButton class="flex-col w-full" type="submit">Продолжить</UButton>
				</UForm>

				<div class="text-sm text-center text-muted mt-2">
					Мы отправим ссылку для сброса пароля на указанный email
				</div>
			</div>
		</UPageCard>
	</div>
</template>
