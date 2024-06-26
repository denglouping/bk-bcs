<template>
  <div class="btn-wrap">
    <bk-select
      v-model="minorDimension"
      :popover-options="{ theme: 'light bk-select-popover add-dimensionality-popover', placement: 'bottom-start' }"
      :popover-min-width="353"
      :search-placeholder="$t('请输入关键字')"
      filterable
      multiple
      @toggle="emits('toggleShowBtn', $event)"
      @change="handleSelectedDimensionChange">
      <template #trigger>
        <span
          v-if="needDown"
          class="action-icon bk-bscp-icon icon-configuration-line"
          v-bk-tooltips="{ content: t('设置维度') }" />
      </template>
      <template #extension>
        <span class="title">{{ t('展示方式') }}</span>
        <bk-radio-group v-model="chartShowType">
          <bk-radio label="tile">
            <span class="bk-bscp-icon icon-bar-chart-thin" />
            <span>{{ t('平铺') }}</span>
          </bk-radio>
          <bk-radio label="pile" :disabled="minorDimension.length > 2">
            <span
              class="bk-bscp-icon icon-bar-chart-stack"
              v-bk-tooltips="{
                content: t('三个以上维度不支持堆叠展示效果'),
                disabled: minorDimension.length <= 2,
                boundary: 'parent',
              }" />
            <span>{{ t('堆叠') }} </span>
          </bk-radio>
        </bk-radio-group>
      </template>
      <bk-option v-for="item in allLabel" :id="item" :key="item" :name="item" :disabled="item === primaryDimension" />
    </bk-select>
    <bk-select
      v-model="downDimension"
      ref="downDimensionSelectRef"
      :popover-options="{ theme: 'light bk-select-popover', placement: 'bottom-start' }"
      :popover-min-width="238"
      :filterable="true"
      :search-placeholder="$t('请输入关键字')"
      @toggle="emits('toggleShowBtn', $event)"
      @change="handleSelectedDownDimension">
      <template #trigger>
        <span
          v-if="needDown"
          class="action-icon bk-bscp-icon icon-download"
          v-bk-tooltips="{ content: t('设置下钻') }" />
      </template>
      <bk-option :id="''">{{ $t('无') }}</bk-option>
      <bk-option
        v-for="item in allLabel"
        :id="item"
        :key="item"
        :name="item"
        :disabled="item === primaryDimension || minorDimension.includes(item)" />
    </bk-select>
    <left-turn-line class="action-icon" @click="emits('refresh')" />
    <FilliscreenLine v-if="!isOpenFullScreen" class="action-icon" @click="handleOpenFullScreen" />
    <UnfullScreen v-else class="action-icon" @click="handleCloseFullScreen" />
  </div>
</template>

<script lang="ts" setup>
  import { ref, watch } from 'vue';
  import { LeftTurnLine, FilliscreenLine, UnfullScreen } from 'bkui-vue/lib/icon';
  import BkMessage from 'bkui-vue/lib/message';
  import { useI18n } from 'vue-i18n';

  const { t } = useI18n();

  const props = defineProps<{
    isOpenFullScreen: boolean;
    allLabel?: string[];
    primaryDimension?: string;
    needDown?: boolean;
    selectDimension?: string[];
    drillDimension?: string;
    isStack?: boolean;
  }>();

  const emits = defineEmits([
    'refresh',
    'toggleFullScreen',
    'toggleShowBtn',
    'selectDimension',
    'toggleChartShowType',
    'selectDownDimension',
  ]);

  const chartShowType = ref(props.isStack ? 'pile' : 'tile');
  const minorDimension = ref(props.selectDimension || []);
  const downDimension = ref(props.drillDimension || '');
  const downDimensionSelectRef = ref();

  watch(
    () => chartShowType.value,
    (val) => {
      emits('toggleChartShowType', val);
    },
  );

  watch(
    () => props.selectDimension,
    () => {
      minorDimension.value = props.selectDimension || [];
    },
  );

  // 打开全屏
  const handleOpenFullScreen = () => {
    emits('toggleFullScreen');
    window.addEventListener('keydown', handleEscClose, { once: true });
    BkMessage({
      theme: 'primary',
      message: t('按 Esc 即可退出全屏模式'),
    });
  };

  const handleCloseFullScreen = () => {
    emits('toggleFullScreen');
    window.removeEventListener('keydown', handleEscClose);
  };

  // Esc按键事件处理
  const handleEscClose = (event: KeyboardEvent) => {
    if (event.code === 'Escape') {
      emits('toggleFullScreen');
    }
  };

  // 设置副维度
  const handleSelectedDimensionChange = () => {
    if (minorDimension.value.length > 2) {
      chartShowType.value = 'tile';
    }
    if (minorDimension.value.includes(downDimension.value)) {
      downDimension.value = '';
      emits('selectDownDimension', '');
    }
    emits('selectDimension', minorDimension.value);
  };

  const handleSelectedDownDimension = (val: string) => {
    emits('selectDownDimension', val);
  };
</script>

<style scoped lang="scss">
  .btn-wrap {
    display: flex;
    .action-icon {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-left: 4px;
      width: 28px;
      height: 28px;
      background: #fafbfd;
      border: 1px solid #dcdee5;
      border-radius: 2px;
      color: #979ba5;
      cursor: pointer;
      &:hover {
        color: #3a84ff;
      }
    }
  }
</style>

<style lang="scss">
  .add-dimensionality-popover {
    .bk-select-search-wrapper {
      width: 220px;
    }
    .bk-select-content {
      display: flex;
      height: 200px;
      .bk-select-dropdown {
        min-width: 240px;
        .is-disabled {
          color: #c4c6cc !important;
        }
      }
      .bk-select-extension {
        position: absolute;
        right: 0;
        top: -1px;
        width: 112px;
        flex-direction: column;
        align-items: normal !important;
        background: #f5f7fa;
        height: 100% !important;
        padding: 8px 0 0 15px;
        color: #63656e;
        .bk-radio-group {
          margin-top: 8px;
          flex-direction: column;
          gap: 12px;
          .bk-radio {
            margin-left: 0;
          }
          .bk-bscp-icon {
            font-size: 16px;
            margin-right: 3px;
          }
        }
      }
    }
  }
</style>
