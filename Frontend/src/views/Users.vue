<template>
    <div class="mt-8 flex flex-col h-full">
        <div class="w-full flex justify-end mb-4">
            <button
              class="bg-blueGray-800 text-white active:bg-blueGray-600 text-sm font-bold uppercase px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150 flex items-center"
              type="button"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#ffffff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="icon icon-tabler icons-tabler-outline icon-tabler-plus mr-2">
                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                <path d="M12 5l0 14" />
                <path d="M5 12l14 0" />
              </svg>
              Novo Utilizador
            </button>
        </div>

        <div class="w-full flex-1">
            <ag-grid-vue
              class="ag-theme-quartz w-full"
              style="height: 25rem;"
              :columnDefs="columnDefs"
              :rowData="rowData"
              :defaultColDef="defaultColDef"
              :gridOptions="gridOptions"
            ></ag-grid-vue>
        </div>
    </div>
</template>

<script>
export default {}
</script>

<style>
@import '~ag-grid-community/styles/ag-grid.css';
@import '~ag-grid-community/styles/ag-theme-quartz.css';
</style>