<!-- {% if Identity %} -->
<!-- Header -->
<div class="flex justify-between items-center py-8">
  <!-- Logo -->
  <img src="./img/Logotype.svg" alt="" />

  <!-- User Profile -->
  <div class="h-max flex items-center relative group">
    <div class="font-medium">
      <!-- {{ Identity.User.Name }} {{ Identity.Device.IP }} -->
    </div>
    <svg
      class="inline ml-2 group-hover:-rotate-90 transition-all"
      xmlns="http://www.w3.org/2000/svg"
      width="16"
      height="16"
      viewBox="0 0 16 16"
      fill="none"
    >
      <path
        d="M2.66656 5.63415L7.01323 9.98082C7.52656 10.4942 8.36656 10.4942 8.8799 9.98082L13.2266 5.63416"
        stroke="#0D0D0D"
        stroke-width="1.5"
        stroke-miterlimit="10"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>

    <!-- Profile Details -->
    <div
      class="pt-3 absolute top-full invisible opacity-0 group-hover:opacity-100 group-hover:visible transition-all"
    >
      <div
        class="w-48 h-44 p-4 bg-white text-sm rounded-xl shadow-sm flex-col justify-start items-start gap-3 flex"
      >
        <div
          class="self-stretch h-10 flex-col justify-start items-start gap-1 flex"
        >
          <div class="text-zinc-500 leading-tight">Компания</div>
          <div class="text-stone-950 leading-tight">
            <!-- {{ Identity.Company.Name }} -->
          </div>
        </div>
        <div
          class="self-stretch h-10 flex-col justify-start items-start gap-1 flex"
        >
          <div class="text-zinc-500 leading-tight">Отдел</div>
          <div class="text-stone-950 leading-tight">
            <!-- {{ Identity.Branch.Name }} -->
          </div>
        </div>
        <div
          class="self-stretch h-10 flex-col justify-start items-start gap-1 flex"
        >
          <div class="text-zinc-500 leading-tight">Номер телефона</div>
          <div class="text-stone-950 leading-tight">
            <!-- {{ Identity.User.Phone }} -->
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<!-- {% else %} -->
<!-- <div class="flex py-8"> -->
<!--   <img src="./img/Logotype.svg" alt="" /> -->
<!-- </div> -->
<!-- {% endif %} -->
