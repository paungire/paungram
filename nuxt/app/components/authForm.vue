<script>
import * as z from "zod";
z.config(z.locales.ru());

export default {
	data() {
		return {
			error: null,
			showPass: false,
			remember: false,

			schema: z.object({
				email: z.string().email(),
				password: z.string().min(8),
			}),

			state: {
				email: "",
				password: "",
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
				this.error = "Ошибка входа, напишите на почту paungire@bk.ru";
			}
		},
	},
};
</script>

<template>
	<div class="p-4">
		<UPageCard class="w-full max-w-md">
			<div>
				<div class="flex flex-col text-center items-center">
					<div class="mb-2">
						<UIcon name="pepicons-print:persons" class="size-12"></UIcon>
					</div>
					<div class="text-xl text-pretty font-semibold text-highlighted">
						Авторизация
					</div>
					<div class="mt-1 text-base text-pretty text-muted">
						Нет аккаунта?
						<ULink to="/register/" class="text-primary font-medium"
							>Пройти регистрацию</ULink
						>.
					</div>
				</div>

				<AuthProvidersList />

				<UForm
					:schema="schema"
					:state="state"
					class="space-y-4"
					@submit="onSubmit"
				>
					<UFormField name="email" label="Email" required>
						<UInput
							v-model="state.email"
							placeholder="Введите email"
							class="w-full"
						/>
					</UFormField>

					<UFormField name="password" label="Пароль" required>
						<template #hint>
							<ULink to="/forgot/" class="text-primary font-medium"
								>Забыли пароль?</ULink
							>
						</template>
						<UInput
							class="w-full"
							v-model="state.password"
							placeholder="Введите пароль"
							:type="showPass ? 'text' : 'password'"
							:ui="{ trailing: 'pe-1' }"
						>
							<template #trailing>
								<UButton
									color="neutral"
									variant="link"
									size="sm"
									:icon="showPass ? 'i-lucide-eye-off' : 'i-lucide-eye'"
									:aria-label="showPass ? 'скрыть' : 'показать'"
									:aria-pressed="showPass"
									aria-controls="password"
									@click="showPass = !showPass"
								/>
							</template>
						</UInput>
					</UFormField>

					<UCheckbox v-model="remember" label="Запомнить меня" />

					<UAlert
						v-if="error"
						color="error"
						icon="i-lucide-info"
						:title="error"
					/>

					<UButton class="flex-col w-full" type="submit">Продолжить</UButton>
				</UForm>

				<div class="text-sm text-center text-muted mt-2">
					Авторизуясь, вы доверяете сервису свои данные и соглашаетесь с
					<ULink to="/policy/" class="text-primary font-medium">правилами</ULink
					>.
				</div>
			</div>
		</UPageCard>
	</div>
</template>
