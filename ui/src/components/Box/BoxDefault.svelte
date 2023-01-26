<div class="border border-gray-200 shadow-lg rounded w-full min-h-15 lg:w-96 xl:w-96">
    <div class="flex justify-center p-1">
        <slot name="box-header"/>
    </div>
    <hr class="border-t border-gray-200"/>
    <div class="px-4 mt-4">
        <slot name="box-body"/>
    </div>
    <div class="px-4">
        <slot name="box-footer"/>
    </div>
</div>