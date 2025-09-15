<script setup lang="ts">
import { computed, ref, watchEffect } from "vue";

const show = ref(false);

// внешний v-model
const password = defineModel<string>("password", { default: "" });

function checkStrength(str: string) {
	const requirements = [
		{ regex: /.{8,}/, text: "Не менее 8 символов" },
		{ regex: /\d/, text: "Не менее 1 цифры" },
		{ regex: /[a-z]/, text: "Не менее 1 строчной буквы" },
		{ regex: /[A-Z]/, text: "Не менее 1 заглавной буквы" },
		{
			regex: /[!@#$%^&*(),.?":{}|<>]/,
			text: "Не менее 1 специального символа",
		},
	];

	return requirements.map((req) => ({
		met: req.regex.test(str),
		text: req.text,
	}));
}

const strength = computed(() => checkStrength(password.value));
const score = computed(() => strength.value.filter((req) => req.met).length);

const color = computed(() => {
	if (score.value === 0 && !colorize.value) return "neutral";
	if (score.value <= 2) return "error";
	if (score.value <= 4) return "warning";
	return "success";
});

const colorize = ref(false);

watchEffect(() => {
	if (score.value === 5) {
		colorize.value = false;
	} else if (score.value > 0) {
		colorize.value = true;
	}
});

const text = computed(() => {
	if (score.value === 0) return "Введите пароль";
	if (score.value <= 2) return "Слабый пароль";
	if (score.value <= 4) return "Средний пароль";
	return "Надежный пароль";
});

function onBlurToColorize() {
	colorize.value = true;
}

defineExpose({ score, text });
</script>

<template>
	<div>
		<UFormField label="Пароль" required>
			<div class="space-y-2">
				<UInput
					v-model="password"
					placeholder="Введите пароль"
					:color="color"
					:type="show ? 'text' : 'password'"
					:aria-invalid="score < 4"
					aria-describedby="password-strength"
					:ui="{ trailing: 'pe-1' }"
					class="w-full"
					:class="
						colorize ? `rounded-md ring-accented ring-1 ring-${color}` : ''
					"
					@blur="onBlurToColorize()"
				>
					<template #trailing>
						<UButton
							color="neutral"
							variant="link"
							size="sm"
							:icon="show ? 'i-lucide-eye-off' : 'i-lucide-eye'"
							:aria-label="show ? 'Скрыть' : 'Показать'"
							:aria-pressed="show"
							aria-controls="password"
							@click="show = !show"
						/>
					</template>
				</UInput>

				<UProgress
					:color="color"
					:indicator="text"
					:model-value="score"
					:max="5"
					size="sm"
				/>
			</div>

			<div v-if="colorize" class="mt-2" :class="`text-${color}`">
				{{ text }}
			</div>
		</UFormField>

		<ul class="space-y-1 mt-2" aria-label="Password requirements">
			<li
				v-for="(req, index) in strength"
				:key="index"
				class="flex items-center gap-0.5"
				:class="req.met ? 'text-success' : 'text-muted'"
			>
				<UIcon
					:name="req.met ? 'i-lucide-circle-check' : 'i-lucide-circle-x'"
					class="size-4 shrink-0"
				/>

				<span class="text-xs font-light">
					{{ req.text }}
					<span class="sr-only">
						{{ req.met ? " - Requirement met" : " - Requirement not met" }}
					</span>
				</span>
			</li>
		</ul>
	</div>
</template>
