<script>
    import TaskTable from '$lib/TaskTable.svelte';
    import TaskTableRow from '$lib/TaskTableRow.svelte';
    import PrimaryButton from '$lib/UI/PrimaryButton.svelte';

    const tasks = [
      {
    "id": "1008924",
    "number": 8924,
    "user_id": "0b49ecb3-b68b-482a-9fae-b493bc0be39e",
    "name": "Test",
    "subject": "<h2>Development Environment Company от TestUser:</h2>\n<h3>Суть обращения:</h3>\n<p>Test Subject</p>\n<hr>\n<h3>Дополнительная информания:</h3>\n<ul>\n<li>Контакты: 61928391923</li>\n<li>Устройство: 127.0.0.1 (127.0.0.1/32)</li>\n<li>Отдел:<br/>\nНазвание: Dev Branch<br/>\nОписание: Master Branch<br/>\nАдрес: Development Space<br/>\nКонтакты: dev@env.com</li>\n</ul>",
    "status": "assigned",
    "timeCreated": {
        "value": "2023-10-09T05:54:54+00:00"
    }
},{
    "id": "1008924",
    "number": 8924,
    "user_id": "0b49ecb3-b68b-482a-9fae-b493bc0be39e",
    "name": "Test",
    "subject": "<h2>Development Environment Company от TestUser:</h2>\n<h3>Суть обращения:</h3>\n<p>Test Subject</p>\n<hr>\n<h3>Дополнительная информания:</h3>\n<ul>\n<li>Контакты: 61928391923</li>\n<li>Устройство: 127.0.0.1 (127.0.0.1/32)</li>\n<li>Отдел:<br/>\nНазвание: Dev Branch<br/>\nОписание: Master Branch<br/>\nАдрес: Development Space<br/>\nКонтакты: dev@env.com</li>\n</ul>",
    "status": "assigned",
    "timeCreated": {
        "value": "2023-10-09T05:54:54+00:00"
    }
},{
    "id": "1008924",
    "number": 8924,
    "user_id": "0b49ecb3-b68b-482a-9fae-b493bc0be39e",
    "name": "Test",
    "subject": "<h2>Development Environment Company от TestUser:</h2>\n<h3>Суть обращения:</h3>\n<p>Test Subject</p>\n<hr>\n<h3>Дополнительная информания:</h3>\n<ul>\n<li>Контакты: 61928391923</li>\n<li>Устройство: 127.0.0.1 (127.0.0.1/32)</li>\n<li>Отдел:<br/>\nНазвание: Dev Branch<br/>\nОписание: Master Branch<br/>\nАдрес: Development Space<br/>\nКонтакты: dev@env.com</li>\n</ul>",
    "status": "assigned",
    "timeCreated": {
        "value": "2023-10-09T05:54:54+00:00"
    }
}
    ]
</script>
<div class="flex justify-between items-center mt-16 mb-10">
  <h1 class="text-4xl font-bold">Обращения</h1>
  <PrimaryButton>
    Новое обращение
  </PrimaryButton>
</div>
<TaskTable>
  {#each tasks as task}
  <TaskTableRow task={task} />
  {/each}

</TaskTable>
