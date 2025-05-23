<script lang="ts">
	import { ChatService, EditorService, type Project, type TableList } from '$lib/services';
	import Menu from '$lib/components/navbar/Menu.svelte';
	import { Table } from 'lucide-svelte';
	import { getLayout } from '$lib/context/layout.svelte';
	import { overflowToolTip } from '$lib/actions/overflow';

	interface Props {
		project: Project;
	}

	async function loadTables() {
		tables = await ChatService.listTables(project.assistantID, project.id);
	}

	let { project }: Props = $props();

	let menu: ReturnType<typeof Menu>;
	let tables: TableList | undefined = $state();
	const layout = getLayout();
</script>

<Menu
	bind:this={menu}
	title="Tables"
	description="Click to view or edit table data"
	onLoad={loadTables}
>
	{#snippet icon()}
		<Table class="h-5 w-5" />
	{/snippet}
	{#snippet body()}
		{#if !tables}
			<p class="text-gray pt-6 pb-3 text-center text-sm dark:text-gray-300">Loading...</p>
		{:else if !tables.tables || tables.tables.length === 0}
			<p class="text-gray pt-6 pb-3 text-center text-sm dark:text-gray-300">No tables</p>
		{:else}
			<ul class="space-y-4 px-3 py-6 text-sm">
				{#each tables.tables as table}
					<li class="group">
						<div class="flex">
							<button
								class="flex max-w-full flex-1 items-center truncate"
								onclick={async () => {
									await EditorService.load(layout.items, project, 'table://' + table.name);
									layout.fileEditorOpen = true;
									menu?.toggle(false);
								}}
							>
								<Table class="size-5 min-w-fit" />
								<span class="ms-2 group-hover:underline" use:overflowToolTip>{table.name}</span>
							</button>
						</div>
					</li>
				{/each}
			</ul>
		{/if}
	{/snippet}
</Menu>
