<?= $this->extend('layouts/main') ?>


<?= $this->section('content') ?>

<div class="card my-4">
    <div class="card-body">
        <h2>Hostname: <?= $hostname ?></h2>
    </div>
</div>

<div class="card my-4">
    <div class="card-body">
        <h2>Address: <?= $address ?> (Local address: <?= $local_address ?>)</h2>
        <h2>Port: <?= $port ?></h2>
    </div>
</div>


<div class="card my-4">
    <div class="card-body">
        <h2>MAC Address: <?= $MAC ?></h2>
    </div>
</div>


<?= $this->endSection() ?>