<script>
import { z } from "zod";
z.config(z.locales.ru());

export default {
	data() {
		return {
			state: {
				email: "",
				password: "",
			},
			error: null,

			schema: z.object({
				email: z.string().email(),
				password: z.string().superRefine((val, ctx) => {
					if (this.$refs.passwordRef.score < 5) {
						ctx.addIssue({
							code: z.ZodIssueCode.custom,
							message: this.$refs.passwordRef.text,
						});
					}
				}),
			}),
		};
	},
	methods: {
		onSubmit(payload) {
			if (this.$refs.passwordRef.score < 5) {
				this.$refs.formRef.value.setErrors({ password: "Обязательное поле" });
				return;
			}

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
						<UIcon name="pepicons-print:person-plus" class="size-12"></UIcon>
					</div>
					<div class="text-xl text-pretty font-semibold text-highlighted">
						Регистрация
					</div>
					<div class="mt-1 text-base text-pretty text-muted">
						Есть аккаунт?
						<ULink to="/auth/" class="text-primary font-medium">Войти</ULink>.
					</div>
				</div>

				<AuthProvidersList />

				<UForm
					ref="formRef"
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

					<passRegister
						ref="passwordRef"
						:password="state.password"
						v-model="state.password"
					></passRegister>

					<UAlert
						v-if="error"
						color="error"
						icon="i-lucide-info"
						:title="error"
					/>

					<UButton class="flex-col w-full" type="submit">Продолжить</UButton>
				</UForm>

				<div class="text-sm text-center text-muted mt-2">
					Регистрируясь, вы доверяете сервису свои данные и соглашаетесь с
					<ULink to="/policy/" class="text-primary font-medium">правилами</ULink
					>.
				</div>
			</div>
		</UPageCard>
	</div>
</template>
